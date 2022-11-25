package chessdotcom

import (
	"fmt"
	"testing"
)

const (
	player = "erik"
)

func TestPlayerProfile(t *testing.T) {

	client := New()

	t.Run("Should return player profile", func(t *testing.T) {

		profile, err := client.PlayerProfile(player)
		if err != nil {
			t.Fatalf("Failed to fetch player profile: %v", err)
		}

		// PlayerId will always be static and unique
		if profile.PlayerId != 41 {
			t.Fatal("Incorrect PlayerId returned from PlayerProfile")
		}
	})
}

func TestPlayerStats(t *testing.T) {

	client := New()

	t.Run("Should return player stats", func(t *testing.T) {

		stats, err := client.PlayerStats(player)
		if err != nil {
			t.Fatalf("Failed to fetch player stats: %v", err)
		}

		if stats.Fide < 1707 {
			t.Fatal("Incorrect FIDE returned from PlayerStats")
		}

		if stats.ChessDaily.Best.Rating < 2065 {
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

		gameArchives, err := client.PlayerGameArchives(player)
		if err != nil {
			t.Fatalf("Failed to fetch player game archives: %v", err)
		}

		if len(gameArchives) == 0 {
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

		clubs, err := client.PlayerClubs(player)
		if err != nil {
			t.Fatalf("Failed to fetch player clubs: %v", err)
		}

		fmt.Println("Club name", clubs[0].Name)
	})
}

func TestPlayerTournaments(t *testing.T) {

	client := New()

	t.Run("Should return player tournaments", func(t *testing.T) {

		tournaments, err := client.PlayerTournaments(player)
		if err != nil {
			t.Fatalf("Failed to fetch player tournaments: %v", err)
		}

		fmt.Println("Finished status", tournaments.Finished[0].Status)

	})
}

func TestTitledPlayers(t *testing.T) {

	client := New()

	t.Run("Should return titled players", func(t *testing.T) {

		titledPlayers, err := client.TitledPlayers()
		if err != nil {
			t.Fatalf("Failed to fetch titled players: %v", err)
		}

		fmt.Println("Titled 0", titledPlayers[0])
		fmt.Println("Titled 1", titledPlayers[1])
		fmt.Println("Titled 2", titledPlayers[2])
	})
}
