package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type getHostsResponseType struct {
	Data []string
}

func (c *Client) GetHosts() ([]string, error) {
	var getHostsResponse getHostsResponseType

	// Parse the URL:
	parsedURL, err := c.GetBaseHost()
	if err != nil {
		return nil, err
	}

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
	err = json.NewDecoder(res.Body).Decode(&getHostsResponse)
	if err != nil {
		return nil, err
	}

	// Select a host:
	return getHostsResponse.Data, nil
}
