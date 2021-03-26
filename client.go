package audiusclient

import "sync"

type Client struct {
	mu          sync.Mutex
	appName     string
	currentHost string
}

func NewClient(appName string) *Client {
	return &Client{
		appName: appName,
	}
}
