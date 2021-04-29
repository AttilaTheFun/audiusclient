package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetTrendingTracksResponse struct {
	Data []APITrack
}

func (c *Client) GetTrendingTracks(genre string, time string) (GetTrendingTracksResponse, error) {
	var getTrendingTracksResponse GetTrendingTracksResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return getTrendingTracksResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/tracks/trending"

	// Build the query:
	values := requestURL.Query()
	if genre != "" {
		values.Set("genre", genre)
	}
	if time != "" {
		values.Set("time", time)
	}
	requestURL.RawQuery = values.Encode()

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getTrendingTracksResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getTrendingTracksResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getTrendingTracksResponse, err
		}
		err = errors.New(string(body))

		return getTrendingTracksResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getTrendingTracksResponse)
	if err != nil {
		return getTrendingTracksResponse, err
	}

	// Set the stream urls on all of the tracks:
	for index, track := range getTrendingTracksResponse.Data {
		streamURL := *selectedHostURL
		streamURL.Path = "/v1/tracks/" + track.ID + "/stream"
		track.StreamURL = streamURL.String()
		getTrendingTracksResponse.Data[index] = track
	}

	return getTrendingTracksResponse, nil
}
