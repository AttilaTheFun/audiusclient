package audiusclient

// HostFetcher -

type HostFetcher interface {
	FetchHosts() ([]string, error)
}
