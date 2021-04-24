package audiusclient

import (
	"log"
	"testing"
)

func TestHostSelectionServiceSelectHost(t *testing.T) {
	hostFetcher := NewDiscoveryHostFetcher("audiusclient")
	hostHealthCheckFetcher := NewDiscoveryHostHealthCheckFetcher("audiusclient")
	config := NewHostSelectionServiceConfig()
	service := NewHostSelectionService(hostFetcher, hostHealthCheckFetcher, config)

	// Select the host
	selectedHost, err := service.GetSelectedHost()
	if err != nil {
		t.Fatalf("Failed to select host with error: %v", err.Error())
	}

	log.Printf("Selected host: %v", selectedHost)
	log.Println()
}
