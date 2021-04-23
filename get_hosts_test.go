package audiusclient

import (
	"log"
	"testing"
)

func TestGetHosts(t *testing.T) {
	client := NewClient("audiusclient")
	hosts, err := client.GetHosts()
	if err != nil {
		t.Fatalf("Failed to select host with error: %v", err.Error())
	}

	log.Printf("Fetched hosts: %v", hosts)
	log.Println()
}
