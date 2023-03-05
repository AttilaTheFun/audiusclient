package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SearchUsersResponse struct {
	Data []APIUser
}

func (c *Client) SearchUsers(query string) (SearchUsersResponse, error) {
	var searchUsersResponse SearchUsersResponse

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return searchUsersResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/users/search"

	// Build the query:
	values := requestURL.Query()
	values.Set("query", query)
	requestURL.RawQuery = values.Encode()

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return searchUsersResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return searchUsersResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return searchUsersResponse, err
		}
		err = errors.New(string(body))

		return searchUsersResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&searchUsersResponse)
	if err != nil {
		return searchUsersResponse, err
	}

	return searchUsersResponse, nil
}
