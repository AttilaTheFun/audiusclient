package audiusclient

// import (
// 	"errors"
// 	"strings"
// )

// func (c *Client) SelectHost() error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	// Fetch all of the potential hosts:
// 	potentialHosts, err := c.GetHosts()
// 	if err != nil {
// 		return err
// 	}
// 	if len(potentialHosts) == 0 {
// 		return errors.New("audius has no available hosts")
// 	}

// 	// First look for an "official" audius host if we can find one.
// 	selectedHost := matchingSuffix(potentialHosts, "audius.co")
// 	if selectedHost == "" {
// 		// Then look for a "staked" host if that failed.
// 		selectedHost = matchingSubstring(potentialHosts, "staked")
// 	}
// 	if selectedHost == "" {
// 		// Finally just fall back on the first host if all else fails.
// 		selectedHost = potentialHosts[0]
// 	}
// 	c.currentHost = selectedHost

// 	return nil
// }

// func matchingSuffix(strs []string, suffix string) string {
// 	for _, str := range strs {
// 		if strings.HasSuffix(str, suffix) {
// 			return str
// 		}
// 	}

// 	return ""
// }

// func matchingSubstring(strs []string, substring string) string {
// 	for _, str := range strs {
// 		if strings.Contains(str, substring) {
// 			return str
// 		}
// 	}

// 	return ""
// }
