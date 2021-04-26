package audiusclient

import (
	"log"
	"testing"
)

func TestHostHealthCheckServiceSelectHost(t *testing.T) {
	fetcher := newCreatorHostHealthCheckFetcher("audiusclient")
	service := NewHostHealthCheckService(fetcher)

	// Health check the hosts:
	resultMap := service.HealthCheckHosts([]string{"https://creatornode2.audius.co"})

	log.Printf("Health checked hosts: %+v", resultMap)
	log.Println()
}
