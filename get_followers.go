package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type GetFollowersResponse struct {
	Data []APIUser
}

func (c *Client) GetFollowers(userID string, limit uint64, offset uint64) (GetFollowersResponse, error) {
	var getFollowersResponse GetFollowersResponse

	// Build the query:
	query := url.Values{}
	query.Add("limit", strconv.FormatUint(limit, 10))
	query.Add("offset", strconv.FormatUint(offset, 10))
	query.Add("user_id", userID)

	// Select an audius host:
	selectedHostURL, err := c.GetHost()
	if err != nil {
		return getFollowersResponse, err
	}
	requestURL := *selectedHostURL
	requestURL.Path = "/v1/full/users/" + userID + "/followers"
	requestURL.RawQuery = query.Encode()

	// Create the request:
	urlString := requestURL.String()
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return getFollowersResponse, err
	}

	// Make the request:
	res, err := httpClient.Do(req)
	if err != nil {
		return getFollowersResponse, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return getFollowersResponse, err
		}
		err = errors.New(string(body))

		return getFollowersResponse, err
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getFollowersResponse)
	if err != nil {
		return getFollowersResponse, err
	}

	return getFollowersResponse, nil
}
