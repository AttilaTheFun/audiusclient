package audiusclient

import (
	"log"
	"testing"
)

func TestCreatorHostHealthCheckFetcher(t *testing.T) {
	hostHealthCheckFetcher := NewCreatorHostHealthCheckFetcher("audiusclient")
	duration, err := hostHealthCheckFetcher.FetchHostHealthCheck("https://creatornode.audius.co")
	if err != nil {
		t.Fatalf("Failed to health check host with error: %v", err.Error())
	}

	log.Printf("Health checked host in: %v", duration)
	log.Println()
}
