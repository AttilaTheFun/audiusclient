package audiusclient

import "time"

type hostHealthCheckServiceConfig struct {

	// How long to wait before health checking a host again. Defaults to one hour.
	HostHealthCheckResultTTL time.Duration
}

func newHostHealthCheckServiceConfig() hostHealthCheckServiceConfig {
	return hostHealthCheckServiceConfig{
		HostHealthCheckResultTTL: time.Hour,
	}
}
