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

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getTrackResponse, err
	}
	parsedURL.Path = "/v1/tracks/" + trackID

	// Create the request:
	urlString := parsedURL.String()
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

	return getTrackResponse, nil
}
