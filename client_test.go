package chessdotcom

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	client := New()

	ppr, err := client.PlayerProfile("erik")
	if err != nil {
		t.Fatalf("Failed to fetch player profile: %v", err)
	}

	fmt.Printf("%+vn", ppr)
}
