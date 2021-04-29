package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SearchPlaylistsResponse struct {
	Data []APIPlaylist
}

func (c *Client) SearchPlaylists(query string) (SearchPlaylistsResponse, error) {
	var searchPlaylistsResponse SearchPlaylistsResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return searchPlaylistsResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/playlists/search"

	// Build the query:
	values := requestURL.Query()
	values.Set("query", query)
	requestURL.RawQuery = values.Encode()

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return searchPlaylistsResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return searchPlaylistsResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return searchPlaylistsResponse, err
		}
		err = errors.New(string(body))

		return searchPlaylistsResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&searchPlaylistsResponse)
	if err != nil {
		return searchPlaylistsResponse, err
	}

	return searchPlaylistsResponse, nil
}
