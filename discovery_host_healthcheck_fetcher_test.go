package audiusclient

import (
	"log"
	"testing"
)

func TestDiscoveryHostHealthCheckFetcher(t *testing.T) {
	hostHealthCheckFetcher := newDiscoveryHostHealthCheckFetcher("audiusclient")
	duration, err := hostHealthCheckFetcher.FetchHostHealthCheck("https://discoveryprovider.audius2.prod-us-west-2.staked.cloud")
	if err != nil {
		t.Fatalf("Failed to health check host with error: %v", err.Error())
	}

	log.Printf("Health checked host in: %v", duration)
	log.Println()
}
