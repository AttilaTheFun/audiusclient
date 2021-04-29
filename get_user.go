package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetUserResponse struct {
	Data APIUser
}

func (c *Client) GetUser(userID string) (GetUserResponse, error) {
	var getUserResponse GetUserResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return getUserResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/users/" + userID

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getUserResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getUserResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getUserResponse, err
		}
		err = errors.New(string(body))

		return getUserResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getUserResponse)
	if err != nil {
		return getUserResponse, err
	}

	return getUserResponse, nil
}
