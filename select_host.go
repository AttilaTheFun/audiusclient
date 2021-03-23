package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type selectHostResponseType struct {
	Data []string
}

func (c *Client) SelectHost() error {
	var getHostsResponse selectHostResponseType

	// Parse the URL:
	parsedURL, err := c.GetBaseHost()
	if err != nil {
		return err
	}

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := http.Get(urlString)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// Parse the error:
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	// Decode the body:
	err = json.NewDecoder(res.Body).Decode(&getHostsResponse)
	if err != nil {
		return err
	}

	// Select a host:
	if len(getHostsResponse.Data) == 0 {
		return errors.New("Unable to retrieve Audius host")
	}
	c.currentHost = getHostsResponse.Data[0]

	return nil
}
