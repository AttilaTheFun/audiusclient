package audiusclient

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	client := NewClient("audiusclient")
	getUserResponse, err := client.GetUser("n0AML")
	if err != nil {
		t.Fatalf("Failed to get playlist tracks with error: %v", err.Error())
	}

	t.Logf("Get playlist tracks response: %v", getUserResponse)
	log.Println(getUserResponse.Data.ID)
	log.Println()
}
