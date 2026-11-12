package audiusclient

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

type empty struct{}

func TestGetFolloweesThatAreNotFollowers(t *testing.T) {

	// Setup:
	service := NewHostSelectionService("audiusclient")
	client := NewClient(service)
	userID := "n0AML"
	limit := 50

	// Get the user:
	getFullUserResponse, err := client.GetFullUser(userID)
	if err != nil {
		t.Fatalf("Failed to get user with error: %v", err.Error())
	}
	followeeCount := getFullUserResponse.Data[0].FolloweeCount
	followerCount := getFullUserResponse.Data[0].FollowerCount

	log.Printf("Followee count: %v", followeeCount)
	log.Printf("Follower count: %v", followerCount)

	log.Printf("Fetching all followee ids")

	// Get *all* of the followee IDs:
	followeeBatches := (followeeCount / limit) + 1
	followeeBatch := 0
	var followeeIDs []string
	for followeeBatch < followeeBatches {
		getFolloweesResponse, err := client.GetFollowees(userID, uint64(limit), uint64(followeeBatch*limit))
		if err != nil {
			log.Printf("Failed to get followees with error: %v", err.Error())
			continue
		}

		log.Printf("Fetched followee batch: %v / %v", followeeBatch+1, followeeBatches-1)

		for _, user := range getFolloweesResponse.Data {
			followeeIDs = append(followeeIDs, user.ID)
		}

		followeeBatch += 1
	}

	log.Printf("All followee ids: %v", followeeIDs)

	log.Printf("Fetching all follower ids")

	// Get *all* of the follower IDs:
	followerBatches := (followerCount / limit) + 1
	followerBatch := 0
	var followerIDs []string
	for followerBatch < followerBatches {
		getFollowersResponse, err := client.GetFollowers(userID, uint64(limit), uint64(followerBatch*limit))
		if err != nil {
			log.Printf("Failed to get followers with error: %v", err.Error())
			continue
		}

		log.Printf("Fetched follower batch: %v / %v", followerBatch+1, followerBatches-1)

		for _, user := range getFollowersResponse.Data {
			followerIDs = append(followerIDs, user.ID)
		}

		followerBatch += 1
	}

	log.Printf("All follower ids: %v", followerIDs)

	// Create a map of all follower IDs:
	followerIDMap := map[string]empty{}
	for _, followerID := range followerIDs {
		followerIDMap[followerID] = empty{}
	}

	// Make a list of all the followees that are not followers:
	var followeeIDsThatAreNotFollowers []string
	for _, followeeID := range followeeIDs {
		_, ok := followerIDMap[followeeID]
		if !ok {
			followeeIDsThatAreNotFollowers = append(followeeIDsThatAreNotFollowers, followeeID)
		}
	}

	followeeIDsThatAreNotFollowersCount := len(followeeIDsThatAreNotFollowers)
	log.Printf("Followees that are not followers count: %v", followeeIDsThatAreNotFollowersCount) // 1748

	// Create a CSV file for writing:
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	dir := filepath.Join(homeDir, "Downloads")
	filePath := filepath.Join(dir, "followees_that_are_not_followers.csv")
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	// Write the results to a CSV file:
	f.WriteString("User ID,Link\n")

	// Print out the links:
	for _, followeeID := range followeeIDsThatAreNotFollowers {
		_, err := f.WriteString(followeeID + ",https://audius.co/users/" + followeeID + "\n")
		if err != nil {
			panic(err)
		}
	}

	log.Println()
}
