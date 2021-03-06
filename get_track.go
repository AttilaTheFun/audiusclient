package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetTrackResponse struct {
	Data APITrack
}

func (c *Client) GetTrack(trackID string) (GetTrackResponse, error) {
	var getTrackResponse GetTrackResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return getTrackResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/tracks/" + trackID

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getTrackResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getTrackResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getTrackResponse, err
		}
		err = errors.New(string(body))

		return getTrackResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getTrackResponse)
	if err != nil {
		return getTrackResponse, err
	}

	// Set the stream urls on all of the tracks:
	streamURL := *selectedHostURL
	streamURL.Path = "/v1/tracks/" + getTrackResponse.Data.ID + "/stream"
	track := getTrackResponse.Data
	track.StreamURL = streamURL.String()
	getTrackResponse.Data = track

	return getTrackResponse, nil
}
