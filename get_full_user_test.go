package audiusclient

import (
	"log"
	"testing"
)

func TestGetFullUser(t *testing.T) {
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	getFullUserResponse, err := client.GetFullUser("n0AML")
	if err != nil {
		t.Fatalf("Failed to get user with error: %v", err.Error())
	}

	t.Logf("Get full user response: %v", getFullUserResponse)
	log.Println(getFullUserResponse.Data[0].ID)
	log.Println()
}
