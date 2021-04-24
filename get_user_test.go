package audiusclient

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	client := NewClient("audiusclient")
	getUserResponse, err := client.GetUser("n0AML")
	if err != nil {
		t.Fatalf("Failed to get user with error: %v", err.Error())
	}

	t.Logf("Get user response: %v", getUserResponse)
	log.Println(getUserResponse.Data.ID)
	log.Println()
}
