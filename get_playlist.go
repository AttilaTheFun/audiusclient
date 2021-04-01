package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetPlaylistResponse struct {
	Data APIPlaylist
}

type internalGetPlaylistResponse struct {

	// TODO: Make this an individual track once Audius fixes the endpoint.
	Data []APIPlaylist
}

func (c *Client) GetPlaylist(playlistID string) (GetPlaylistResponse, error) {
	var getPlaylistResponse GetPlaylistResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getPlaylistResponse, err
	}
	parsedURL.Path = "/v1/playlists/" + playlistID

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
	if err != nil {
		return getPlaylistResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getPlaylistResponse, err
		}
		err = errors.New(string(body))

		return getPlaylistResponse, err
	}

	// Decode the body:
	var internalResponse internalGetPlaylistResponse
	err = json.NewDecoder(res.Body).Decode(&internalResponse)
	if err != nil {
		return getPlaylistResponse, err
	}
	if len(internalResponse.Data) == 0 {
		return getPlaylistResponse, errors.New("unable to find the playlist for the given ID")
	}
	getPlaylistResponse.Data = internalResponse.Data[0]

	return getPlaylistResponse, nil
}
