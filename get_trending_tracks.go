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
	var getTrendingTracks GetTrendingTracksResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getTrendingTracks, err
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

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
	if err != nil {
		return getTrendingTracks, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getTrendingTracks, err
		}
		err = errors.New(string(body))

		return getTrendingTracks, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getTrendingTracks)
	if err != nil {
		return getTrendingTracks, err
	}

	return getTrendingTracks, nil
}
