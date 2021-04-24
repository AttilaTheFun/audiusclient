package audiusclient

type Client struct {
	hostSelectionService *HostSelectionService
}

func NewClient(hostSelectionService *HostSelectionService) *Client {
	return &Client{
		hostSelectionService: hostSelectionService,
	}
}
