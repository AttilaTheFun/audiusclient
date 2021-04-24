package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DiscoveryHostFetcher struct {
	appName string
}

func NewDiscoveryHostFetcher(appName string) *DiscoveryHostFetcher {
	return &DiscoveryHostFetcher{
		appName: appName,
	}
}

type fetchDiscoveryHostsResponseType struct {
	Data []string
}

func (f DiscoveryHostFetcher) FetchHosts() ([]string, error) {

	// Parse the url:
	parsedURL, err := url.Parse("https://api.audius.co")
	if err != nil {
		return nil, err
	}

	// Add the query params:
	values := url.Values{}
	values.Set("app_name", f.appName)
	parsedURL.RawQuery = values.Encode()

	// Create the request:
	urlString := parsedURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {

		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(string(body))
	}

	// Decode the body:
	var fetchDiscoveryHostsResponse fetchDiscoveryHostsResponseType
	err = json.NewDecoder(res.Body).Decode(&fetchDiscoveryHostsResponse)
	if err != nil {
		return nil, err
	}

	// Select a host:
	return fetchDiscoveryHostsResponse.Data, nil
}
