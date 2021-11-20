package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetFullUserResponse struct {
	Data []APIUser
}

func (c *Client) GetFullUser(userID string) (GetFullUserResponse, error) {
	var getFullUserResponse GetFullUserResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return getFullUserResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/full/users/" + userID

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getFullUserResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getFullUserResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getFullUserResponse, err
		}
		err = errors.New(string(body))

		return getFullUserResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getFullUserResponse)
	if err != nil {
		return getFullUserResponse, err
	}

	return getFullUserResponse, nil
}
