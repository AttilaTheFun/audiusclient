package audiusclient

type hostFetcher interface {
	FetchHosts() ([]string, error)
}
