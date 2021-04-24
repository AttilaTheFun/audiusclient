package audiusclient

import "net/url"

// func (c *Client) GetBaseHost() (*url.URL, error) {

// 	// Parse the host:
// 	parsedHost, err := url.Parse("https://api.audius.co")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Add the query params:
// 	values := url.Values{}
// 	values.Set("app_name", c.appName)
// 	parsedHost.RawQuery = values.Encode()

// 	return parsedHost, nil
// }

func (c *Client) GetHost() (*url.URL, error) {

	// Get the selected host:
	selectedHost, err := c.hostSelectionService.GetSelectedHost()
	if err != nil {
		return nil, err
	}

	// Parse the host:
	parsedHost, err := url.Parse(selectedHost)
	if err != nil {
		return nil, err
	}

	// Add the query params:
	values := url.Values{}
	values.Set("app_name", c.hostSelectionService.appName)
	parsedHost.RawQuery = values.Encode()

	return parsedHost, nil
}
