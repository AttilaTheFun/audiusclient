package audiusclient

import (
	"log"
	"testing"
)

func TestGetTrendingTracks(t *testing.T) {
	client := NewClient("audiusclient")
	getTrendingTracks, err := client.GetTrendingTracks("Electronic", "month")
	if err != nil {
		t.Fatalf("Failed to get trending tracks with error: %v", err.Error())
	}

	t.Logf("Get trending tracks response: %v", getTrendingTracks)
	log.Printf("%+v", getTrendingTracks.Data[0].ID)
	log.Println()
}