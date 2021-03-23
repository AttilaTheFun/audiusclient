package audiusclient

func (c *Client) GetTrackStream(trackID string) (string, error) {

	// Parse the Audius host url:
	parsedURL, err := c.GetHost()
	if err != nil {
		return "", err
	}
	parsedURL.Path = "/v1/tracks/" + trackID + "/stream"

	// Create the url string:
	urlString := parsedURL.String()

	return urlString, nil
}
