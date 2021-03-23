package audiusclient

import "net/url"

func (c *Client) GetBaseHost() (*url.URL, error) {

	// Parse the host:
	parsedHost, err := url.Parse("https://api.audius.co")
	if err != nil {
		return nil, err
	}

	// Add the query params:
	values := url.Values{}
	values.Set("app_name", "audiusclient")
	parsedHost.RawQuery = values.Encode()

	return parsedHost, nil
}

func (c *Client) GetHost() (*url.URL, error) {

	// Select a host if necessary:
	if c.currentHost == "" {
		err := c.SelectHost()
		if err != nil {
			return nil, err
		}
	}

	// Parse the host:
	parsedHost, err := url.Parse(c.currentHost)
	if err != nil {
		return nil, err
	}

	// Add the query params:
	values := url.Values{}
	values.Set("app_name", "audiusclient")
	parsedHost.RawQuery = values.Encode()

	return parsedHost, nil
}
