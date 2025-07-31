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

	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New().WithField("module", "qrlClients")

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

type QRLClients struct {
	ClientReleaseVersion string
	ClientReleaseDate    template.HTML
	IsUserSubscribed     bool
}

type QRLClientServicesPageData struct {
	LastUpdate time.Time
	Gzond      QRLClients
	Qrysm      QRLClients
	CsrfField  template.HTML
}

var qrlClients = QRLClientServicesPageData{}
var qrlClientsMux = &sync.RWMutex{}

var httpClient = &http.Client{Timeout: time.Second * 10}

// Init starts a go routine to update the QRL Clients Info
func Init() {
	go update()
}

func fetchClientData(repo string) *gitAPIResponse {
	var gitAPI = new(gitAPIResponse)

	githubAPIHost := utils.Config.GithubApiHost
	if githubAPIHost == "" {
		githubAPIHost = "api.github.com"
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://%s/repos%s/releases/latest", githubAPIHost, repo))

	if err != nil {
		logger.Errorf("error retrieving QRL Client Data: %v", err)
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("error retrieving QRL Client Data, status code: %v", resp.StatusCode)
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(&gitAPI)

	if err != nil {
		logger.Errorf("error decoding QRL Clients json response to struct: %v", err)
		return nil
	}

	return gitAPI
}

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

func prepareQRLClientData(repo string) (string, template.HTML) {
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

		return client.Name, utils.FormatTimestamp(rTime.Unix())
	}
	return "Github", "searching" // If API limit is exceeded
}

func updateQRLClient() {
	curTime := time.Now()
	// sending 8 requests to github per call
	// git api rate-limit 60 per hour : 60/8 = 7.5 minutes minimum
	if curTime.Sub(qrlClients.LastUpdate) < time.Hour { // LastUpdate is initialized at January 1, year 1 so no need to check for nil
		return
	}

	logger.Println("Updating QRL Clients Information")
	qrlClientsMux.Lock()
	defer qrlClientsMux.Unlock()
	qrlClients.Gzond.ClientReleaseVersion, qrlClients.Gzond.ClientReleaseDate = prepareQRLClientData("/theQRL/go-zond")

	qrlClients.Qrysm.ClientReleaseVersion, qrlClients.Qrysm.ClientReleaseDate = prepareQRLClientData("/theQRL/qrysm")

	qrlClients.LastUpdate = curTime
}

func update() {
	for {
		updateQRLClient()
		time.Sleep(time.Minute * 5)
	}
}

// GetQRLClientData returns a QRLClientServicesPageData
func GetQRLClientData() QRLClientServicesPageData {
	qrlClientsMux.Lock()
	defer qrlClientsMux.Unlock()
	return qrlClients
}
