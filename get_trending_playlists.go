package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetTrendingPlaylistsResponse struct {
	Data []APIPlaylist
}

func (c *Client) GetTrendingPlaylists(time string) (GetTrendingPlaylistsResponse, error) {
	var getTrendingPlaylistsResponse GetTrendingPlaylistsResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getTrendingPlaylistsResponse, err
	}
	parsedURL.Path = "/v1/playlists/trending"

	// Build the query:
	values := parsedURL.Query()
	if time != "" {
		values.Set("time", time)
	}
	parsedURL.RawQuery = values.Encode()

	// Create the request:
	urlString := parsedURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getTrendingPlaylistsResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getTrendingPlaylistsResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getTrendingPlaylistsResponse, err
		}
		err = errors.New(string(body))

		return getTrendingPlaylistsResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getTrendingPlaylistsResponse)
	if err != nil {
		return getTrendingPlaylistsResponse, err
	}

	return getTrendingPlaylistsResponse, nil
}
