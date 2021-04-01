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
	var getTrendingPlaylists GetTrendingPlaylistsResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getTrendingPlaylists, err
	}
	parsedURL.Path = "/v1/playlists/trending"

	// Build the query:
	values := parsedURL.Query()
	if time != "" {
		values.Set("time", time)
	}
	parsedURL.RawQuery = values.Encode()

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
	if err != nil {
		return getTrendingPlaylists, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getTrendingPlaylists, err
		}
		err = errors.New(string(body))

		return getTrendingPlaylists, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getTrendingPlaylists)
	if err != nil {
		return getTrendingPlaylists, err
	}

	return getTrendingPlaylists, nil
}
