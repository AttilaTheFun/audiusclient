package audiusclient

import (
	"time"
)

// HostHealthCheckFetcher -

type HostHealthCheckFetcher interface {

	// Calls the health check endpoint on the host.
	// Returns the time to first response byte if successful or an error otherwise.
	FetchHostHealthCheck(host string) (time.Duration, error)
}
