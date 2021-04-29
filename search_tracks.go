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

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return searchTracksResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/tracks/search"

	// Build the query:
	values := requestURL.Query()
	values.Set("query", query)
	requestURL.RawQuery = values.Encode()

	// Create the request:
	urlString := requestURL.String()
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

	// Set the stream urls on all of the tracks:
	for index, track := range searchTracksResponse.Data {
		streamURL := *selectedHostURL
		streamURL.Path = "/v1/tracks/" + track.ID + "/stream"
		track.StreamURL = streamURL.String()
		searchTracksResponse.Data[index] = track
	}

	return searchTracksResponse, nil
}
