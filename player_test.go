package chessdotcom

import (
	"testing"
)

const (
	player = "erik"
)

func TestPlayerProfile(t *testing.T) {

	client := New()

	t.Run("Should return player profile", func(t *testing.T) {

		res, err := client.PlayerProfile(player)
		if err != nil {
			t.Fatalf("Failed to fetch player profile: %v", err)
		}

		// PlayerId will always be static and unique
		if res.PlayerId != 41 {
			t.Fatal("Incorrect PlayerId returned from PlayerProfile")
		}
	})
}

func TestPlayerStats(t *testing.T) {

	client := New()

	t.Run("Should return player stats", func(t *testing.T) {

		res, err := client.PlayerStats(player)
		if err != nil {
			t.Fatalf("Failed to fetch player stats: %v", err)
		}

		if res.Fide < 1707 {
			t.Fatal("Incorrect FIDE returned from PlayerStats")
		}

		if res.ChessDaily.Best.Rating < 2065 {
			t.Fatal("Incorrect best daily rating from PlayerStats")
		}
	})
}

func TestPlayerGames(t *testing.T) {

	client := New()

	t.Run("Should return player games", func(t *testing.T) {

		_, err := client.PlayerGames(player)
		if err != nil {
			t.Fatalf("Failed to fetch player games: %v", err)
		}

	})
}

func TestPlayerArchives(t *testing.T) {

	client := New()

	t.Run("Should return player game archives", func(t *testing.T) {

		res, err := client.PlayerGameArchives(player)
		if err != nil {
			t.Fatalf("Failed to fetch player game archives: %v", err)
		}

		if len(res.Archives) == 0 {
			t.Fatal("No archives returned from PlayerGameArchives")
		}
	})
}

func TestPlayerGamesToMove(t *testing.T) {

	client := New()

	t.Run("Should return player games to move", func(t *testing.T) {

		_, err := client.PlayerGamesToMove(player)
		if err != nil {
			t.Fatalf("Failed to fetch player games to move: %v", err)
		}
	})
}

func TestPlayerGameArchivePGN(t *testing.T) {

	client := New()

	t.Run("Should return player game archive PGN", func(t *testing.T) {

		_, err := client.PlayerGameArchivePGN(player, "2022", "01")
		if err != nil {
			t.Fatalf("Failed to fetch player game archives PGN: %v", err)
		}
	})
}

func TestPlayerClubs(t *testing.T) {

	client := New()

	t.Run("Should return player clubs", func(t *testing.T) {

		_, err := client.PlayerClubs(player)
		if err != nil {
			t.Fatalf("Failed to fetch player clubs: %v", err)
		}

	})
}

func TestPlayerTournaments(t *testing.T) {

	client := New()

	t.Run("Should return player tournaments", func(t *testing.T) {

		_, err := client.PlayerTournaments(player)
		if err != nil {
			t.Fatalf("Failed to fetch player tournaments: %v", err)
		}

	})
}

func TestTitledPlayers(t *testing.T) {

	client := New()

	t.Run("Should return titled players", func(t *testing.T) {

		_, err := client.TitledPlayers()
		if err != nil {
			t.Fatalf("Failed to fetch titled players: %v", err)
		}
	})
}
