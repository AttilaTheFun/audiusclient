package audiusclient

import (
	"log"
	"testing"
)

func TestGetResolvedResource(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	resourceType, resourceID, err := client.GetResolvedResource("https://audius.co/wuki/hotline-stay-and-bling-wukileak--425934")
	if err != nil {
		t.Fatalf("Failed to get track stream with error: %v", err.Error())
	}

	log.Printf("Resource type: %v, resourceID: %v", resourceType, resourceID)
	log.Println()
}
