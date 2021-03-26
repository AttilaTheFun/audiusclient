package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrack(t *testing.T) {
	client := NewClient("audiusclient")
	getTrackResponse, err := client.GetTrack("BRG99") // "n3RMe"
	if err != nil {
		t.Fatalf("Failed to get track with error: %v", err.Error())
	}

	t.Logf("Get track response: %v", getTrackResponse)
	log.Println(getTrackResponse.Data.Artwork["1000x1000"])
	log.Println()
}
