package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

type creatorHostHealthCheckFetcher struct {
	appName string
}

func NewCreatorHostHealthCheckFetcher(appName string) *creatorHostHealthCheckFetcher {
	return &creatorHostHealthCheckFetcher{
		appName: appName,
	}
}

type creatorHostHealtcheckData struct {
	ServiceProviderID          int64  `json:"spID"`
	ServiceProviderOwnerWallet string `json:"spOwnerWallet"`
	IsRegisteredOnURSM         bool   `json:"isRegisteredOnURSM"`

	SelectedDiscoveryProvider string `json:"selectedDiscoveryProvider"`
	CreatorNodeEndpoint       string `json:"creatorNodeEndpoint"`

	Git     string `json:"git"`
	Service string `json:"service"`
	Version string `json:"version"`
	Healthy bool   `json:"healthy"`
}

type fetchCreatorHostHealthCheckResponseType struct {
	Data creatorHostHealtcheckData
}

func (f creatorHostHealthCheckFetcher) FetchHostHealthCheck(host string) (time.Duration, error) {

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
	var fetchCreatorHostHealthCheckResponse fetchCreatorHostHealthCheckResponseType
	err = json.NewDecoder(res.Body).Decode(&fetchCreatorHostHealthCheckResponse)
	if err != nil {
		return 0, err
	}

	// Check if the host is healthy:
	if !fetchCreatorHostHealthCheckResponse.Data.Healthy {
		return 0, errors.New("host is unhealthy")
	}

	return duration, nil
}
