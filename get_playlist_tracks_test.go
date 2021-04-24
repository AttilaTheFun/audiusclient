package audiusclient

import (
	"log"
	"testing"
)

func TestGetPlaylistTracks(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getPlaylistTracksResponse, err := client.GetPlaylistTracks("nqZmb")
	if err != nil {
		t.Fatalf("Failed to get playlist tracks with error: %v", err.Error())
	}

	t.Logf("Get playlist tracks response: %v", getPlaylistTracksResponse)
	log.Println(getPlaylistTracksResponse.Data[0].Title)
	log.Println()
}
