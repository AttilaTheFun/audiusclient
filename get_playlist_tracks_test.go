package audiusclient

import (
	"log"
	"testing"
)

func TestGetPlaylistTracks(t *testing.T) {
	client := NewClient("audiusclient")
	getPlaylistTracksResponse, err := client.GetPlaylistTracks("nZaYa")
	if err != nil {
		t.Fatalf("Failed to get playlist tracks with error: %v", err.Error())
	}

	t.Logf("Get playlist tracks response: %v", getPlaylistTracksResponse)
	log.Println(getPlaylistTracksResponse.Data[0].ID)
	log.Println()
}
