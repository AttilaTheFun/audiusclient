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

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getTrendingTracksResponse, err
	}
	parsedURL.Path = "/v1/tracks/trending"

	// Build the query:
	values := parsedURL.Query()
	if genre != "" {
		values.Set("genre", genre)
	}
	if time != "" {
		values.Set("time", time)
	}
	parsedURL.RawQuery = values.Encode()

	// Create the request:
	urlString := parsedURL.String()
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

	return getTrendingTracksResponse, nil
}
