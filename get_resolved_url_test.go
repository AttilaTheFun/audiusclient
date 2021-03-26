package audiusclient

import (
	"log"
	"testing"
)

func TestGetResolvedURL(t *testing.T) {
	client := NewClient("audiusclient")
	resourceType, resourceID, err := client.GetResolvedURL("https://audius.co/audiusplaylists/playlist/lo-fi-nights-4629")
	if err != nil {
		t.Fatalf("Failed to get track stream with error: %v", err.Error())
	}

	log.Printf("Resource type: %v, resourceID: %v", resourceType, resourceID)
	log.Println()
}
