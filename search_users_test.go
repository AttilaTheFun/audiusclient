package audiusclient

import (
	"log"
	"testing"
)

func TestSearchUsers(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	searchUsersResponse, err := client.SearchUsers("ON THE HUNT")
	if err != nil {
		t.Fatalf("Failed to search users with error: %v", err.Error())
	}

	t.Logf("Search users response: %v", searchUsersResponse)
	log.Printf("%+v", searchUsersResponse.Data[0].ProfilePicture)
	log.Println()
	log.Printf("%+v", searchUsersResponse.Data[0].CoverPhoto)
	log.Println()
}
