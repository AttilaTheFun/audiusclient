package audiusclient

type Client struct {
	appName     string
	currentHost string
}

func NewClient(appName string) *Client {
	return &Client{
		appName:     appName,
		currentHost: "",
	}
}
