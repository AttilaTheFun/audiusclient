package audiusclient

import (
	"log"
	"testing"
)

func TestGetFollowees(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getFolloweesResponse, err := client.GetFollowees("n0AML", 50, 0)
	if err != nil {
		t.Fatalf("Failed to get user with error: %v", err.Error())
	}

	t.Logf("Get followees response: %v", getFolloweesResponse)

	// Extract followee IDs:
	var followeeIDs []string
	for _, user := range getFolloweesResponse.Data {
		followeeIDs = append(followeeIDs, user.ID)
	}

	log.Printf("Followee ids: %v", followeeIDs)
	log.Println()
}
