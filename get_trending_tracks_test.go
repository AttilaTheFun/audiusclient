package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrendingTracks(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getTrendingTracks, err := client.GetTrendingTracks("Hip-Hop/Rap", "year")
	if err != nil {
		t.Fatalf("Failed to get trending tracks with error: %v", err.Error())
	}

	t.Logf("Get trending tracks response: %v", getTrendingTracks)
	log.Printf("%+v", getTrendingTracks.Data[0].Title)
	log.Println()
}
