package audiusclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetResolvedURLResponse struct {
	Data APITrack
}

// GetResolvedURL - Resolves an audius url into a resource type (playlist, track, or user) and resource ID.
func (c *Client) GetResolvedURL(audiusURL string) (resourceType string, resourceID string, err error) {

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return
	}
	parsedURL.Path = "/v1/resolve"

	// Build the query:
	values := parsedURL.Query()
	values.Set("url", audiusURL)
	parsedURL.RawQuery = values.Encode()

	// Create a client that won't redirect:
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Fetch the hosts:
	urlString := parsedURL.String()
	res, err := client.Get(urlString)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 302 {
		// Parse the error:
		var body []byte
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}
		err = errors.New(string(body))

		return
	}

	// Get the redirected url:
	redirectedURL, err := res.Location()
	if err != nil {
		return
	}

	// Parse the redirected URL into the resource type and ID:
	resourcePath := strings.TrimPrefix(redirectedURL.Path, "/v1/")
	resourcePathComponents := strings.Split(resourcePath, "/")
	if len(resourcePathComponents) != 2 {
		err = errors.New("invalid resource path")
		return
	}
	resourceID = resourcePathComponents[1]
	switch resourcePathComponents[0] {
	case "playlists":
		resourceType = "playlist"
	case "tracks":
		resourceType = "track"
	case "users":
		resourceType = "user"
	default:
		err = errors.New("invalid resource path")
		return
	}

	return
}
