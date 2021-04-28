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
	var getPlaylistTracksResponse GetPlaylistTracksResponse

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getPlaylistTracksResponse, err
	}
	parsedURL.Path = "/v1/playlists/" + playlistID + "/tracks"

	// Create the request:
	urlString := parsedURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getPlaylistTracksResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getPlaylistTracksResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getPlaylistTracksResponse, err
		}
		err = errors.New(string(body))

		return getPlaylistTracksResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getPlaylistTracksResponse)
	if err != nil {
		return getPlaylistTracksResponse, err
	}

	// Set the stream urls on all of the tracks:
	for _, track := range getPlaylistTracksResponse.Data {
		streamURL := *parsedURL
		streamURL.Path = "/v1/tracks/" + track.ID + "/stream"
		track.StreamURL = streamURL.String()
	}

	return getPlaylistTracksResponse, nil
}
