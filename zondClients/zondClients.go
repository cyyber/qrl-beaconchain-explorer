package zondclients

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New().WithField("module", "zondClients")

type ethernodesAPIStruct struct {
	Client string `json:"client"`
	Value  int    `json:"value"`
}
type gitAPIResponse struct {
	URL           string        `json:"url"`
	AssetsURL     string        `json:"assets_url"`
	UploadURL     string        `json:"upload_url"`
	HTMLURL       string        `json:"html_url"`
	ID            uint64        `json:"id"`
	Author        interface{}   `json:"author"`
	NodeID        string        `json:"node_id"`
	TagName       string        `json:"tag_name"`
	Target        string        `json:"target_commitish"`
	Name          string        `json:"name"`
	Draft         bool          `json:"draft"`
	PreRelease    bool          `json:"prerelease"`
	CreatedDate   string        `json:"created_at"`
	PublishedDate string        `json:"published_at"`
	Assets        []interface{} `json:"assets"`
	Tarball       string        `json:"tarball_url"`
	ZipBall       string        `json:"zipball_url"`
	Body          string        `json:"body"`
}

type clientUpdateInfo struct {
	Name string
	Date time.Time
}

type ZondClients struct {
	ClientReleaseVersion string
	ClientReleaseDate    template.HTML
	NetworkShare         string
	IsUserSubscribed     bool
}

type ZondClientServicesPageData struct {
	LastUpdate time.Time
	Gzond      ZondClients
	Qrysm      ZondClients
	CsrfField  template.HTML
}

var zondClients = ZondClientServicesPageData{}
var zondClientsMux = &sync.RWMutex{}
var bannerClients = []clientUpdateInfo{}
var bannerClientsMux = &sync.RWMutex{}

var httpClient = &http.Client{Timeout: time.Second * 10}

// Init starts a go routine to update the Zond Clients Info
func Init() {
	// TODO(rgeraldes24)
	// go update()
}

func fetchClientData(repo string) *gitAPIResponse {
	var gitAPI = new(gitAPIResponse)

	githubAPIHost := utils.Config.GithubApiHost
	if githubAPIHost == "" {
		githubAPIHost = "api.github.com"
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://%s/repos%s/releases/latest", githubAPIHost, repo))

	if err != nil {
		logger.Errorf("error retrieving ZOND Client Data: %v", err)
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("error retrieving ZOND Client Data, status code: %v", resp.StatusCode)
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(&gitAPI)

	if err != nil {
		logger.Errorf("error decoding ZOND Clients json response to struct: %v", err)
		return nil
	}

	return gitAPI
}

/*
var ethernodesAPI []ethernodesAPIStruct

func fetchClientNetworkShare() []ethernodesAPIStruct {
	resp, err := http.Get("https://ethernodes.org/api/clients")

	if err != nil {
		logger.Errorf("error retrieving ETH Clients Network Share Data: %v", err)
		return ethernodesAPI
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ethernodesAPI)

	if err != nil {
		logger.Errorf("error decoding ETH Clients Network Share json response to struct: %v", err)
	}

	return ethernodesAPI
}
*/

func getRepoTime(date string, dTime string) (time.Time, error) {
	var year, month, day, hour, min int64
	var err error
	dateDays := strings.Split(date, "-")
	dateTimes := strings.Split(dTime, ":")
	if len(dateDays) < 3 || len(dateTimes) < 3 {
		return time.Now(), fmt.Errorf("invalid date string %s %s", date, dTime)
	}
	year, err = strconv.ParseInt(dateDays[0], 10, 0)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing year: %v", err)
	}
	month, err = strconv.ParseInt(dateDays[1], 10, 0)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing month: %v", err)
	}
	day, err = strconv.ParseInt(dateDays[2], 10, 0)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing day: %v", err)
	}
	hour, err = strconv.ParseInt(dateTimes[0], 10, 0)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing hour: %v", err)
	}
	min, err = strconv.ParseInt(dateTimes[1], 10, 0)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing min: %v", err)
	}
	return time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(min), 0, 0, time.UTC), nil
}

func prepareZondClientData(repo string, name string, curTime time.Time) (string, template.HTML) {
	client := fetchClientData(repo)
	time.Sleep(time.Millisecond * 250) // consider github rate limit

	if client == nil {
		return "Github", "searching"
	}
	date := strings.Split(client.PublishedDate, "T")

	if len(date) == 2 {
		rTime, err := getRepoTime(date[0], date[1])
		if err != nil {
			logger.Errorf("error parsing git repo. time: %v", err)
			return client.Name, "GitHub" // client.Name is client version from github api
		}
		timeDiff := (curTime.Sub(rTime).Hours() / 24.0)

		if timeDiff < 1 { // add recent releases for notification collector to be collected
			update := clientUpdateInfo{Name: name, Date: rTime}
			bannerClients = append(bannerClients, update)
		}
		return client.Name, utils.FormatTimestamp(rTime.Unix())
	}
	return "Github", "searching" // If API limit is exceeded
}

func updateZondClientNetShare() {
	/*
		nShare := fetchClientNetworkShare()
		total := 0
		for _, item := range nShare {
			total += item.Value
		}

		for _, item := range nShare {
			share := fmt.Sprintf("%.1f%%", (float64(item.Value)/float64(total))*100.0)
			switch item.Client {
			case "gzond":
				zondClients.Gzond.NetworkShare = share
			default:
				continue
			}
		}
	*/
}

func updateZondClient() {
	curTime := time.Now()
	// sending 8 requests to github per call
	// git api rate-limit 60 per hour : 60/8 = 7.5 minutes minimum
	if curTime.Sub(zondClients.LastUpdate) < time.Hour { // LastUpdate is initialized at January 1, year 1 so no need to check for nil
		return
	}

	logger.Println("Updating Zond Clients Information")
	zondClientsMux.Lock()
	defer zondClientsMux.Unlock()
	bannerClientsMux.Lock()
	defer bannerClientsMux.Unlock()
	bannerClients = []clientUpdateInfo{}
	updateZondClientNetShare()
	zondClients.Gzond.ClientReleaseVersion, zondClients.Gzond.ClientReleaseDate = prepareZondClientData("/theQRL/go-zond", "Gzond", curTime)

	zondClients.Qrysm.ClientReleaseVersion, zondClients.Qrysm.ClientReleaseDate = prepareZondClientData("/theQRL/qrysm", "Qrysm", curTime)

	zondClients.LastUpdate = curTime
}

func update() {
	for {
		updateZondClient()
		time.Sleep(time.Minute * 5)
	}
}

// GetZondClientData returns a ZondClientServicesPageData
func GetZondClientData() ZondClientServicesPageData {
	zondClientsMux.Lock()
	defer zondClientsMux.Unlock()
	return zondClients
}

// GetUpdatedClients returns a slice of latest updated clients or empty slice if no updates
func GetUpdatedClients() []clientUpdateInfo {
	bannerClientsMux.Lock()
	defer bannerClientsMux.Unlock()
	return bannerClients
	// return []string{"Qrysm"}
}
