package audiusclient

import (
	"log"
	"testing"
)

func TestGetPlaylist(t *testing.T) {
	client := NewClient("audiusclient")
	getPlaylistResponse, err := client.GetPlaylist("nZaYa")
	if err != nil {
		t.Fatalf("Failed to get playlist tracks with error: %v", err.Error())
	}

	t.Logf("Get playlist tracks response: %v", getPlaylistResponse)
	log.Println(getPlaylistResponse.Data.User.ID)
	log.Println()
}