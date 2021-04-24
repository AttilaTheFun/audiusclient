package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrendingPlaylists(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getTrendingPlaylists, err := client.GetTrendingPlaylists("week")
	if err != nil {
		t.Fatalf("Failed to get trending playlists with error: %v", err.Error())
	}

	t.Logf("Get trending playlists response: %v", getTrendingPlaylists)
	log.Printf("%+v", getTrendingPlaylists.Data[0].PlaylistName)
	log.Println()
}
