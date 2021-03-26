package audiusclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type selectHostResponseType struct {
	Data []string
}

func (c *Client) SelectHost() error {
	c.mu.Lock()
	defer c.mu.Unlock()

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
	potentialHosts := getHostsResponse.Data
	if len(potentialHosts) == 0 {
		return errors.New("Unable to retrieve Audius host")
	}
	// First look for an "official" audius host if we can find one.
	selectedHost := matchingSuffix(potentialHosts, "audius.co")
	if selectedHost == "" {
		// Then look for a "staked" host if that failed.
		selectedHost = matchingSubstring(potentialHosts, "staked")
	}
	if selectedHost == "" {
		// Finally just fall back on the first host if all else fails.
		selectedHost = potentialHosts[0]
	}
	c.currentHost = selectedHost

	return nil
}

func matchingSuffix(strs []string, suffix string) string {
	for _, str := range strs {
		if strings.HasSuffix(str, suffix) {
			return str
		}
	}

	return ""
}

func matchingSubstring(strs []string, substring string) string {
	for _, str := range strs {
		if strings.Contains(str, substring) {
			return str
		}
	}

	return ""
}
