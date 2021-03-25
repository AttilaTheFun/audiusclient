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

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return getUserResponse, err
	}
	parsedURL.Path = "/v1/users/" + userID

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
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
