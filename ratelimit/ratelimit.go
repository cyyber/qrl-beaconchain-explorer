package ratelimit

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/metrics"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

type TimeWindow string

const (
	SecondTimeWindow = "second"
	HourTimeWindow   = "hour"
	MonthTimeWindow  = "month"

	HeaderRateLimitLimit     = "ratelimit-limit"     // the rate limit ceiling that is applicable for the current request
	HeaderRateLimitRemaining = "ratelimit-remaining" // the number of requests left for the current rate-limit window
	HeaderRateLimitReset     = "ratelimit-reset"     // the number of seconds until the quota resets
	HeaderRateLimitWindow    = "ratelimit-window"    // what window the ratelimit represents
	HeaderRetryAfter         = "retry-after"         // the number of seconds until the quota resets, same as HeaderRateLimitReset, RFC 7231, 7.1.3

	HeaderRateLimitRemainingSecond = "x-ratelimit-remaining-second" // the number of requests left for the current rate-limit window
	HeaderRateLimitRemainingMinute = "x-ratelimit-remaining-minute" // the number of requests left for the current rate-limit window
	HeaderRateLimitRemainingHour   = "x-ratelimit-remaining-hour"   // the number of requests left for the current rate-limit window
	HeaderRateLimitRemainingDay    = "x-ratelimit-remaining-day"    // the number of requests left for the current rate-limit window
	HeaderRateLimitRemainingMonth  = "x-ratelimit-remaining-month"  // the number of requests left for the current rate-limit window

	HeaderRateLimitLimitSecond = "x-ratelimit-limit-second" // the rate limit ceiling that is applicable for the current user
	HeaderRateLimitLimitMinute = "x-ratelimit-limit-minute" // the rate limit ceiling that is applicable for the current user
	HeaderRateLimitLimitHour   = "x-ratelimit-limit-hour"   // the rate limit ceiling that is applicable for the current user
	HeaderRateLimitLimitDay    = "x-ratelimit-limit-day"    // the rate limit ceiling that is applicable for the current user
	HeaderRateLimitLimitMonth  = "x-ratelimit-limit-month"  // the rate limit ceiling that is applicable for the current user

	DefaultRateLimitSecond = 2   // RateLimit per second if no ratelimits are set in database
	DefaultRateLimitHour   = 500 // RateLimit per second if no ratelimits are set in database
	DefaultRateLimitMonth  = 0   // RateLimit per second if no ratelimits are set in database

	FallbackRateLimitSecond = 20 // RateLimit per second for when redis is offline
	FallbackRateLimitBurst  = 20 // RateLimit burst for when redis is offline

	defaultBucket = "default" // if no bucket is set for a route, use this one

	statsTruncateDuration = time.Hour * 1 // ratelimit-stats are truncated to this duration
)

var redisClient *redis.Client
var redisIsHealthy atomic.Bool

var fallbackRateLimiter = NewFallbackRateLimiter() // if redis is offline, use this rate limiter

var initializedWg = &sync.WaitGroup{} // wait for everything to be initialized before serving requests

var logger = logrus.StandardLogger().WithField("module", "ratelimit")

type RateLimit struct {
	Second int64
	Hour   int64
	Month  int64
}

type RateLimitResult struct {
	BlockRequest  bool
	Time          time.Time
	Weight        int64
	Route         string
	IP            string
	Key           string
	RedisKeys     []RedisKey
	RedisStatsKey string
	RateLimit     *RateLimit

	Limit       int64
	LimitSecond int64
	LimitMinute int64
	LimitHour   int64
	LimitDay    int64
	LimitMonth  int64

	Remaining       int64
	RemainingSecond int64
	RemainingMinute int64
	RemainingHour   int64
	RemainingDay    int64
	RemainingMonth  int64

	Reset  int64
	Bucket string
	Window TimeWindow
}

type RedisKey struct {
	Key      string
	ExpireAt time.Time
}

type responseWriterDelegator struct {
	http.ResponseWriter
	written     int64
	status      int
	wroteHeader bool
}

func (r *responseWriterDelegator) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	n, err := r.ResponseWriter.Write(b)
	r.written += int64(n)
	return n, err
}

func (r *responseWriterDelegator) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
	r.wroteHeader = true
}

func (r *responseWriterDelegator) Status() int {
	return r.status
}

var DefaultRequestFilter = func(req *http.Request) bool {
	if req.URL == nil || !strings.HasPrefix(req.URL.Path, "/api") || strings.HasPrefix(req.URL.Path, "/api/i/") || strings.HasPrefix(req.URL.Path, "/api/v1/docs/") || strings.HasPrefix(req.URL.Path, "/api/v2/docs/") {
		return false
	}
	return true
}

var requestFilter = DefaultRequestFilter
var requestFilterMu = &sync.RWMutex{}

func SetRequestFilter(filter func(req *http.Request) bool) {
	requestFilterMu.Lock()
	defer requestFilterMu.Unlock()
	requestFilter = filter
}

func GetRequestFilter() func(req *http.Request) bool {
	requestFilterMu.RLock()
	defer requestFilterMu.RUnlock()
	return requestFilter
}

var maxBadRequestWeight int64 = 1

func SetMaxBadRquestWeight(weight int64) {
	atomic.StoreInt64(&maxBadRequestWeight, weight)
}

func GetMaxBadRquestWeight() int64 {
	return atomic.LoadInt64(&maxBadRequestWeight)
}

// Init initializes the RateLimiting middleware, the rateLimiting middleware will not work without calling Init first. The second parameter is a function the will get called on every request, it will only apply ratelimiting to requests when this func returns true.
func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        utils.Config.RedisSessionStoreEndpoint,
		ReadTimeout: time.Second * 3,
	})

	initializedWg.Add(1)

	go func() {
		firstRun := true
		for {
			err := updateRedisStatus()
			if err != nil {
				logger.WithError(err).Errorf("error checking redis")
				time.Sleep(time.Second * 1)
				continue
			}
			if firstRun {
				initializedWg.Done()
				firstRun = false
			}
			time.Sleep(time.Second * 1)
		}
	}()

	initializedWg.Wait()
}

// HttpMiddleware returns an http.Handler that can be used as middleware to RateLimit requests. If redis is offline, it will use a fallback rate limiter.
func HttpMiddleware(next http.Handler) http.Handler {
	initializedWg.Wait()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := GetRequestFilter()
		if !f(r) {
			next.ServeHTTP(w, r)
			return
		}

		if !redisIsHealthy.Load() {
			metrics.Counter.WithLabelValues("ratelimit_fallback").Inc()
			fallbackRateLimiter.Handle(w, r, next.ServeHTTP)
			return
		}

		rl, err := rateLimitRequest(r)
		if err != nil {
			// just serve the request if there is a problem with getting the rate limit
			logger.WithFields(logrus.Fields{"error": err}).Errorf("error getting rate limit")
			metrics.Errors.WithLabelValues("ratelimit_rateLimitRequest").Inc()
			next.ServeHTTP(w, r)
			return
		}

		// logrus.WithFields(logrus.Fields{"route": rl.Route, "key": rl.Key, "limit": rl.Limit, "remaining": rl.Remaining, "reset": rl.Reset, "window": rl.Window, "validKey": rl.IsValidKey}).Infof("rateLimiting")

		w.Header().Set(HeaderRateLimitLimit, strconv.FormatInt(rl.Limit, 10))
		w.Header().Set(HeaderRateLimitRemaining, strconv.FormatInt(rl.Remaining, 10))
		w.Header().Set(HeaderRateLimitReset, strconv.FormatInt(rl.Reset, 10))

		w.Header().Set(HeaderRateLimitWindow, string(rl.Window))

		w.Header().Set(HeaderRateLimitLimitMonth, strconv.FormatInt(rl.LimitMonth, 10))
		w.Header().Set(HeaderRateLimitLimitDay, strconv.FormatInt(rl.LimitDay, 10))
		w.Header().Set(HeaderRateLimitLimitHour, strconv.FormatInt(rl.LimitHour, 10))
		w.Header().Set(HeaderRateLimitLimitMinute, strconv.FormatInt(rl.LimitMinute, 10))
		w.Header().Set(HeaderRateLimitLimitSecond, strconv.FormatInt(rl.LimitSecond, 10))

		w.Header().Set(HeaderRateLimitRemainingMonth, strconv.FormatInt(rl.RemainingMonth, 10))
		w.Header().Set(HeaderRateLimitRemainingDay, strconv.FormatInt(rl.RemainingDay, 10))
		w.Header().Set(HeaderRateLimitRemainingHour, strconv.FormatInt(rl.RemainingHour, 10))
		w.Header().Set(HeaderRateLimitRemainingMinute, strconv.FormatInt(rl.RemainingMinute, 10))
		w.Header().Set(HeaderRateLimitRemainingSecond, strconv.FormatInt(rl.RemainingSecond, 10))

		if rl.BlockRequest {
			metrics.Counter.WithLabelValues("ratelimit_block").Inc()
			w.Header().Set(HeaderRetryAfter, strconv.FormatInt(rl.Reset, 10))
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			err = postRateLimit(rl, http.StatusTooManyRequests)
			if err != nil {
				logger.WithFields(logrus.Fields{"error": err}).Errorf("error calling postRateLimit")
			}
			return
		}

		d := &responseWriterDelegator{ResponseWriter: w}
		next.ServeHTTP(d, r)
		err = postRateLimit(rl, d.Status())
		if err != nil {
			logger.WithFields(logrus.Fields{"error": err}).Errorf("error calling postRateLimit")
		}
	})
}

// updateRedisStatus checks if redis is healthy and updates redisIsHealthy accordingly.
func updateRedisStatus() error {
	oldStatus := redisIsHealthy.Load()
	newStatus := true
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	defer cancel()
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		logger.WithError(err).Errorf("error pinging redis")
		newStatus = false
	}
	if oldStatus != newStatus {
		logger.WithFields(logrus.Fields{"oldStatus": oldStatus, "newStatus": newStatus}).Infof("redis status changed")
	}
	redisIsHealthy.Store(newStatus)
	return nil
}

// postRateLimit decrements the rate limit keys in redis if the status is not 200.
func postRateLimit(rl *RateLimitResult, status int) error {
	if !(status >= 500 && status <= 599) && status != 429 {
		// any statuscode but 5xx or 429 will count towards the ratelimit
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pipe := redisClient.Pipeline()

	decrByWeight := rl.Weight
	mbrw := GetMaxBadRquestWeight()
	if decrByWeight > mbrw {
		decrByWeight = mbrw
	}

	for _, k := range rl.RedisKeys {
		pipe.DecrBy(ctx, k.Key, decrByWeight)
		pipe.ExpireAt(ctx, k.Key, k.ExpireAt) // make sure all keys have a TTL
	}
	pipe.DecrBy(ctx, rl.RedisStatsKey, 1)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// rateLimitRequest is the main function for rate limiting, it will check the rate limits for the request and update the rate limits in redis.
func rateLimitRequest(r *http.Request) (*RateLimitResult, error) {
	start := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("ratelimit_rateLimitRequest").Observe(time.Since(start).Seconds())
	}()

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*1000)
	defer cancel()

	res := &RateLimitResult{}
	// defer func() { logger.Infof("rateLimitRequest: %+v", *res) }()

	key, ip := getKey(r)
	res.Key = key
	res.IP = ip

	weight, route, bucket := getWeight(r)
	nokeyRatelimit := getDefaultRatelimit()
	res.Weight = weight
	res.Route = route
	res.Bucket = bucket
	res.RateLimit = nokeyRatelimit

	startUtc := start.UTC()
	res.Time = startUtc

	nextHourUtc := time.Now().Truncate(time.Hour).Add(time.Hour)
	nextMonthUtc := time.Date(startUtc.Year(), startUtc.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	timeUntilNextHourUtc := nextHourUtc.Sub(startUtc)
	timeUntilNextMonthUtc := nextMonthUtc.Sub(startUtc)

	normedIP := "ip_" + strings.ReplaceAll(ip, ":", "_")
	rateLimitSecondKey := fmt.Sprintf("rl:c:s:%s:%s", res.Bucket, normedIP)
	rateLimitHourKey := fmt.Sprintf("rl:c:h:%04d-%02d-%02d-%02d:%s:%s", startUtc.Year(), startUtc.Month(), startUtc.Day(), startUtc.Hour(), res.Bucket, normedIP)
	rateLimitMonthKey := fmt.Sprintf("rl:c:m:%04d-%02d:%s:%s", startUtc.Year(), startUtc.Month(), res.Bucket, normedIP)
	statsKey := fmt.Sprintf("rl:s:%04d-%02d-%02d-%02d:%s:%s:%s", startUtc.Year(), startUtc.Month(), startUtc.Day(), startUtc.Hour(), "nokey", res.Route, res.Bucket)
	res.RedisStatsKey = statsKey

	pipe := redisClient.Pipeline()

	var rateLimitSecond, rateLimitHour, rateLimitMonth *redis.IntCmd

	if res.RateLimit.Second > 0 {
		rateLimitSecond = pipe.IncrBy(ctx, rateLimitSecondKey, weight)
		pipe.ExpireNX(ctx, rateLimitSecondKey, time.Second)
	}

	if res.RateLimit.Hour > 0 {
		rateLimitHour = pipe.IncrBy(ctx, rateLimitHourKey, weight)
		pipe.ExpireAt(ctx, rateLimitHourKey, nextHourUtc.Add(time.Second*60)) // expire 1 minute after the window to make sure we do not miss any requests due to time-sync
		res.RedisKeys = append(res.RedisKeys, RedisKey{rateLimitHourKey, nextHourUtc.Add(time.Second * 60)})
	}

	if res.RateLimit.Month > 0 {
		rateLimitMonth = pipe.IncrBy(ctx, rateLimitMonthKey, weight)
		pipe.ExpireAt(ctx, rateLimitMonthKey, nextMonthUtc.Add(time.Second*60)) // expire 1 minute after the window to make sure we do not miss any requests due to time-sync
		res.RedisKeys = append(res.RedisKeys, RedisKey{rateLimitMonthKey, nextMonthUtc.Add(time.Second * 60)})
	}

	pipe.Incr(ctx, statsKey)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	if res.RateLimit.Month > 0 && rateLimitMonth.Val() > res.RateLimit.Month {
		res.Limit = res.RateLimit.Month
		res.Remaining = 0
		res.Reset = int64(timeUntilNextMonthUtc.Seconds())
		res.Window = MonthTimeWindow
		res.BlockRequest = true
	} else if res.RateLimit.Hour > 0 && rateLimitHour.Val() > res.RateLimit.Hour {
		res.Limit = res.RateLimit.Hour
		res.Remaining = 0
		res.Reset = int64(timeUntilNextHourUtc.Seconds())
		res.Window = HourTimeWindow
		res.BlockRequest = true
	} else if res.RateLimit.Second > 0 && rateLimitSecond.Val() > res.RateLimit.Second {
		res.Limit = res.RateLimit.Second
		res.Remaining = 0
		res.Reset = int64(1)
		res.Window = SecondTimeWindow
		res.BlockRequest = true
	} else {
		res.Limit = res.RateLimit.Second
		res.Remaining = res.RateLimit.Second - rateLimitSecond.Val()
		res.Reset = int64(1)
		res.Window = SecondTimeWindow
	}

	if res.RateLimit.Second > 0 {
		res.RemainingSecond = res.RateLimit.Second - rateLimitSecond.Val()
		if res.RemainingSecond < 0 {
			res.RemainingSecond = 0
		}
	}
	if res.RateLimit.Hour > 0 {
		res.RemainingHour = res.RateLimit.Hour - rateLimitHour.Val()
		if res.RemainingHour < 0 {
			res.RemainingHour = 0
		}
	}
	if res.RateLimit.Month > 0 {
		res.RemainingMonth = res.RateLimit.Month - rateLimitMonth.Val()
		if res.RemainingMonth < 0 {
			res.RemainingMonth = 0
		}
	}

	// normalize limit-headers to keep them consistent with previous versions
	if res.RateLimit.Month > 0 {
		res.LimitMonth = res.RateLimit.Month
	} else {
		res.LimitMonth = max(res.RateLimit.Month, res.RateLimit.Hour, res.RateLimit.Second)
		res.RemainingMonth = max(res.RemainingMonth, res.RemainingHour, res.RemainingSecond)
	}
	res.LimitDay = res.LimitMonth
	res.RemainingDay = res.RemainingMonth

	if res.RateLimit.Hour > 0 {
		res.LimitHour = res.RateLimit.Hour
	} else {
		res.LimitHour = res.LimitMonth
		res.RemainingHour = res.RemainingMonth
	}
	res.LimitMinute = res.LimitHour
	res.RemainingMinute = res.RemainingHour

	if res.RateLimit.Second > 0 {
		res.LimitSecond = res.RateLimit.Second
	} else {
		res.LimitSecond = res.LimitHour
	}

	return res, nil
}

func getDefaultRatelimit() (nokeyRatelimit *RateLimit) {
	nokeyRatelimit = &RateLimit{
		Second: DefaultRateLimitSecond,
		Hour:   DefaultRateLimitHour,
		Month:  DefaultRateLimitMonth,
	}
	return nokeyRatelimit
}

func max(vals ...int64) int64 {
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// getKey returns the key used for RateLimiting. It first checks the query params, then the header and finally the ip address.
func getKey(r *http.Request) (key, ip string) {
	ip = getIP(r)
	return "nokey", ip
}

// getWeight returns the weight of an endpoint. if the weight of the endpoint is not defined, it returns 1.
func getWeight(r *http.Request) (cost int64, identifier, bucket string) {
	route := getRoute(r)
	var weight int64 = 1
	bucket = defaultBucket

	return weight, route, bucket
}

func getRoute(r *http.Request) string {
	route := mux.CurrentRoute(r)
	pathTpl, err := route.GetPathTemplate()
	if err != nil {
		return "UNDEFINED"
	}
	return pathTpl
}

// getIP returns the ip address from the http request
func getIP(r *http.Request) string {
	ips := r.Header.Get("CF-Connecting-IP")
	if ips == "" {
		ips = r.Header.Get("X-Forwarded-For")
	}
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String()
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "INVALID"
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1"
		}
		return ip
	}

	return "INVALID"
}

type FallbackRateLimiterClient struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type FallbackRateLimiter struct {
	clients map[string]*FallbackRateLimiterClient
	mu      sync.Mutex
}

func NewFallbackRateLimiter() *FallbackRateLimiter {
	rl := &FallbackRateLimiter{
		clients: make(map[string]*FallbackRateLimiterClient),
	}
	go func() {
		for {
			time.Sleep(time.Minute)
			rl.mu.Lock()
			for ip, client := range rl.clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(rl.clients, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()
	return rl
}

func (rl *FallbackRateLimiter) Handle(w http.ResponseWriter, r *http.Request, next func(writer http.ResponseWriter, request *http.Request)) {
	key, _ := getKey(r)
	rl.mu.Lock()
	if _, found := rl.clients[key]; !found {
		rl.clients[key] = &FallbackRateLimiterClient{limiter: rate.NewLimiter(FallbackRateLimitSecond, FallbackRateLimitBurst)}
	}
	rl.clients[key].lastSeen = time.Now()
	if !rl.clients[key].limiter.Allow() {
		rl.mu.Unlock()
		w.Header().Set(HeaderRateLimitLimit, strconv.FormatInt(FallbackRateLimitSecond, 10))
		w.Header().Set(HeaderRateLimitReset, strconv.FormatInt(1, 10))
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		return
	}
	rl.mu.Unlock()
	next(w, r)
}
