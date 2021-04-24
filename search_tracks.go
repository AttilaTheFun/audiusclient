package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SearchTracksResponse struct {
	Data []APITrack
}

func (c *Client) SearchTracks(query string) (SearchTracksResponse, error) {
	var searchTracksResponse SearchTracksResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return searchTracksResponse, err
	}
	parsedURL.Path = "/v1/tracks/search"

	// Build the query:
	values := parsedURL.Query()
	values.Set("query", query)
	parsedURL.RawQuery = values.Encode()

	// Create the request:
	urlString := parsedURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return searchTracksResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return searchTracksResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return searchTracksResponse, err
		}
		err = errors.New(string(body))

		return searchTracksResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&searchTracksResponse)
	if err != nil {
		return searchTracksResponse, err
	}

	return searchTracksResponse, nil
}
