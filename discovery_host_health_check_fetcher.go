package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

type discoveryHostHealthCheckFetcher struct {
	appName string
}

func NewDiscoveryHostHealthCheckFetcher(appName string) *discoveryHostHealthCheckFetcher {
	return &discoveryHostHealthCheckFetcher{
		appName: appName,
	}
}

type discoveryHostHealtcheckWeb struct {
	BlockHash   string `json:"blockhash"`
	BlockNumber int64  `json:"blocknumber"`
}

type discoveryHostHealtcheckDB struct {
	BlockHash   string `json:"blockhash"`
	BlockNumber int64  `json:"number"`
}

type discoveryHostHealtcheckData struct {
	BlockDifference               int64 `json:"block_difference"`
	MaximumHealthyBlockDifference int64 `json:"maximum_healthy_block_difference"`

	TrendingPlaylistsAgeSec int64                      `json:"trending_playlists_age_sec"`
	TrendingTracksAgeSec    int64                      `json:"trending_tracks_age_sec"`
	Web                     discoveryHostHealtcheckWeb `json:"web"`

	DatabaseConnections int64                     `json:"database_connections"`
	DatabaseSize        int64                     `json:"database_size"`
	Database            discoveryHostHealtcheckDB `json:"db"`

	FilesystemSize int64 `json:"filesystem_size"`
	FilesystemUsed int64 `json:"filesystem_used"`

	NumberOfCPUs int64 `json:"number_of_cpus"`

	TotalMemory      int64 `json:"total_memory"`
	UsedMemory       int64 `json:"used_memory"`
	RedisTotalMemory int64 `json:"redis_total_memory"`

	TransferredBytesPerSecond float64 `json:"transferred_bytes_per_sec"`
	ReceivedBytesPerSecond    float64 `json:"received_bytes_per_sec"`

	Git     string `json:"git"`
	Service string `json:"discovery-node"`
	Version string `json:"version"`
}

type fetchDiscoveryHostHealthCheckResponseType struct {
	Data discoveryHostHealtcheckData
}

func (f discoveryHostHealthCheckFetcher) FetchHostHealthCheck(host string) (time.Duration, error) {

	// Create the request:
	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		return 0, err
	}

	// Add the health check path:
	req.URL.Path = "/health_check"

	// Add the trace to the request:
	startTime := time.Now()
	var endTime time.Time
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			endTime = time.Now()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {

		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}
		err = errors.New(string(body))

		return 0, err
	}

	// Calculate the time it took to receive the first byte:
	duration := endTime.Sub(startTime)

	// Decode the body:
	var fetchDiscoveryHostHealthCheckResponse fetchDiscoveryHostHealthCheckResponseType
	err = json.NewDecoder(res.Body).Decode(&fetchDiscoveryHostHealthCheckResponse)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
