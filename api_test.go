package chessdotcom

import (
	"testing"
)

func TestPlayerProfile(t *testing.T) {

	client := New()

	t.Run("Should return player profile", func(t *testing.T) {

		res, err := client.PlayerProfile("magnuscarlsen")
		if err != nil {
			t.Fatalf("Failed to fetch player profile: %v", err)
		}

		// PlayerId will always be static and unique
		if res.PlayerId != 3889224 {
			t.Fatal("Incorrect PlayerId returned from PlayerProfile")
		}
	})
}

func TestPlayerStats(t *testing.T) {

	client := New()

	t.Run("Should return player stats", func(t *testing.T) {

		res, err := client.PlayerStats("magnuscarlsen")
		if err != nil {
			t.Fatalf("Failed to fetch player stats: %v", err)
		}

		// Theres a problem if Magnus is below 2000. Right, Hans...?
		if res.Fide < 2000 {
			t.Fatal("Incorrect FIDE returned from PlayerStats")
		}

		// Cannot drop below a best rating, so we verify against that
		if res.ChessRapid.Best.Rating < 2862 {
			t.Fatal("Incorrect best rapid rating from PlayerStats")
		}
	})
}
