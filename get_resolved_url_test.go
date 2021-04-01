package audiusclient

import (
	"log"
	"testing"
)

func TestGetResolvedURL(t *testing.T) {
	client := NewClient("audiusclient")
	resourceType, resourceID, err := client.GetResolvedURL("https://audius.co/teendaze/four-more-years-313549")
	if err != nil {
		t.Fatalf("Failed to get track stream with error: %v", err.Error())
	}

	log.Printf("Resource type: %v, resourceID: %v", resourceType, resourceID)
	log.Println()
}
