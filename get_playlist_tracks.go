package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetPlaylistTracksResponse struct {
	Data []APITrack
}

func (c *Client) GetPlaylistTracks(playlistID string) (GetPlaylistTracksResponse, error) {
	var getPlaylistTracks GetPlaylistTracksResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getPlaylistTracks, err
	}
	parsedURL.Path = "/v1/playlists/" + playlistID + "/tracks"

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
	if err != nil {
		return getPlaylistTracks, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getPlaylistTracks, err
		}
		err = errors.New(string(body))

		return getPlaylistTracks, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getPlaylistTracks)
	if err != nil {
		return getPlaylistTracks, err
	}

	return getPlaylistTracks, nil
}
