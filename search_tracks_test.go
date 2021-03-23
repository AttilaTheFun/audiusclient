package audiusclient

import (
	"log"
	"testing"
)

func TestSearchTracks(t *testing.T) {
	client := NewClient("audiusclient")
	searchTracksResponse, err := client.SearchTracks("ON THE HUNT")
	if err != nil {
		t.Fatalf("Failed to search tracks with error: %v", err.Error())
	}

	t.Logf("Search tracks response: %v", searchTracksResponse)
	log.Println(searchTracksResponse.Data[0].ID)
	log.Println()
}
