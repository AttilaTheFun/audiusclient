package audiusclient

import (
	"log"
	"testing"
)

func TestDiscoveryHostFetcher(t *testing.T) {
	hostFetcher := NewDiscoveryHostFetcher("audiusclient")
	hosts, err := hostFetcher.FetchHosts()
	if err != nil {
		t.Fatalf("Failed to fetch discovery hosts with error: %v", err.Error())
	}

	log.Printf("Fetched discovery hosts: %v", hosts)
	log.Println()
}
