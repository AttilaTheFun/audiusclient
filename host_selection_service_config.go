package audiusclient

import "time"

type hostSelectionServiceConfig struct {

	// How long the host list is considered fresh for. Defaults to one day.
	HostListTTL time.Duration

	// How long hosts that failed a health check are blacklisted for. Defaults to one hour.
	UnhealthyHostTTL time.Duration

	// The maximum number of concurrent requests to make when evaluating hosts. Defaults to 10.
	MaximumConcurrentRequests int

	// How often the selected host is re-evaluated. Defaults to 10 minutes.
	SelectedHostTTL time.Duration
}

func newHostSelectionServiceConfig() hostSelectionServiceConfig {
	return hostSelectionServiceConfig{
		HostListTTL:               time.Hour * 24,
		UnhealthyHostTTL:          time.Hour,
		MaximumConcurrentRequests: 10,
		SelectedHostTTL:           time.Minute * 10,
	}
}
