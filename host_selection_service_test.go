package audiusclient

import (
	"log"
	"testing"
)

func TestHostSelectionServiceSelectHost(t *testing.T) {
	service := NewHostSelectionService("audiusclient")

	// Select the host
	selectedHost, err := service.GetSelectedHost()
	if err != nil {
		t.Fatalf("Failed to select host with error: %v", err.Error())
	}

	log.Printf("Selected host: %v", selectedHost)
	log.Println()
}
