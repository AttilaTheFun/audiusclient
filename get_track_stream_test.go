package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrackStream(t *testing.T) {
	client := NewClient("audiusclient")
	streamURL, err := client.GetTrackStream("n3RMe")
	if err != nil {
		t.Fatalf("Failed to get track stream with error: %v", err.Error())
	}

	t.Logf("Track stream: %v", streamURL)
	log.Println(streamURL)
	log.Println()
}
