package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrack(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getTrackResponse, err := client.GetTrack("2PEEp") // "n3RMe"
	if err != nil {
		t.Fatalf("Failed to get track with error: %v", err.Error())
	}

	t.Logf("Get track response: %v", getTrackResponse)
	log.Println(getTrackResponse.Data.StreamURL)
	log.Println()
}
