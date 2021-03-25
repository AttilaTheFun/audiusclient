package audiusclient

import (
	"log"
	"testing"
)

func TestSearchPlaylists(t *testing.T) {
	client := NewClient("audiusclient")
	searchPlaylistsResponse, err := client.SearchPlaylists("trap")
	if err != nil {
		t.Fatalf("Failed to search playlists with error: %v", err.Error())
	}

	t.Logf("Search playlists response: %v", searchPlaylistsResponse)
	log.Println(searchPlaylistsResponse.Data[1].User.ProfilePicture)
	log.Println()
}
