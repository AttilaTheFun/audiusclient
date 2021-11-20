package audiusclient

import (
	"log"
	"testing"
)

func TestGetFollowers(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getFollowersResponse, err := client.GetFollowers("n0AML", 50, 0)
	if err != nil {
		t.Fatalf("Failed to get user with error: %v", err.Error())
	}

	t.Logf("Get followers response: %v", getFollowersResponse)

	// Extract follower IDs:
	var followerIDs []string
	for _, user := range getFollowersResponse.Data {
		followerIDs = append(followerIDs, user.ID)
	}

	log.Printf("Follower ids: %v", followerIDs)
	log.Println()
}
