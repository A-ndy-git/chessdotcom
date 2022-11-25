package chessdotcom

import (
	"fmt"
	"time"
)

type PlayerProfile struct {
	Avatar     string        `json:"avatar"`
	PlayerId   int64         `json:"player_id"`
	ID         string        `json:"id"`
	URL        string        `json:"url"`
	Name       string        `json:"name"`
	Username   string        `json:"username"`
	Followers  int64         `json:"followers"`
	Country    string        `json:"country"`
	Location   string        `json:"location"`
	LastOnline time.Duration `json:"last_online"`
	Joined     time.Duration `json:"joined"`
	Status     string        `json:"status"`
	IsStreamer bool          `json:"is_streamer"`
	Verified   bool          `json:"verified"`
}

// Returns details of an existing player.
func (c *Client) PlayerProfile(username string) (PlayerProfile, error) {
	var res PlayerProfile

	endpoint := fmt.Sprintf("/player/%v", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Last struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
	Rd     int64         `json:"rd"`
}

type Best struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
	Game   string        `json:"game"`
}

type Record struct {
	Win            int64         `json:"win"`
	Loss           int64         `json:"loss"`
	Draw           int64         `json:"draw"`
	TimePerMove    time.Duration `json:"time_per_move"`
	TimeoutPercent int           `json:"timeout_percent"`
}

type Tournament struct {
	Points        int64 `json:"points"`
	Withdraw      int64 `json:"withdraw"`
	Count         int64 `json:"count"`
	HighestFinish int64 `json:"highest_finish"`
}

type ChessDaily struct {
	Last       Last       `json:"last"`
	Best       Best       `json:"best"`
	Record     Record     `json:"record"`
	Tournament Tournament `json:"tournament"`
}

type ChessRapid struct {
	Last   Last   `json:"last"`
	Best   Best   `json:"best"`
	Record Record `json:"record"`
}

type ChessRating struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
}

type Tactics struct {
	Highest ChessRating `json:"highest"`
	Lowest  ChessRating `json:"lowest"`
}

type PuzzleRushBest struct {
	TotalAttempts int64 `json:"total_attempts"`
	Score         int64 `json:"score"`
}

type PlayerStats struct {
	ChessDaily    ChessDaily     `json:"chess_daily"`
	Chess960Daily ChessDaily     `json:"chess960_daily"`
	ChessRapid    ChessRapid     `json:"chess_rapid"`
	ChessBullet   ChessRapid     `json:"chess_bullet"`
	ChessBlitz    ChessRapid     `json:"chess_blitz"`
	Fide          int64          `json:"fide"`
	Tactics       Tactics        `json:"tactics"`
	PuzzleRush    PuzzleRushBest `json:"puzzle_rush"`
}

// Returns ratings, win/loss, and other stats about a players game play, tactics, lessons and puzzle rush score.
func (c *Client) PlayerStats(username string) (PlayerStats, error) {
	var res PlayerStats

	endpoint := fmt.Sprintf("/player/%v/stats", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type PlayerGames struct {
	URL          string        `json:"url"`
	MoveBy       int64         `json:"move_by"`
	PGN          string        `json:"pgn"`
	TimeControl  string        `json:"time_control"`
	LastActivity time.Duration `json:"last_activity"`
	Rated        bool          `json:"rated"`
	Turn         string        `json:"turn"`
	FEN          string        `json:"fen"`
	StartTime    time.Duration `json:"start_time"`
	TimeClass    string        `json:"time_class"`
	Rules        string        `json:"rules"`
	White        string        `json:"white"`
	Black        string        `json:"black"`
}

// Returns an array of daily chess games that a player is currently playing.
func (c *Client) PlayerGames(username string) ([]PlayerGames, error) {
	var res map[string][]PlayerGames

	endpoint := fmt.Sprintf("/player/%v/games", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return nil, err
	}

	return res["games"], nil
}

// Returns an array of monthly game archives available for a player.
func (c *Client) PlayerGameArchives(username string) ([]string, error) {
	var res map[string][]string

	endpoint := fmt.Sprintf("/player/%v/games/archives", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return nil, err
	}

	return res["archives"], nil
}

type PlayerGamesToMove struct {
	URL          string        `json:"url"`
	MoveBy       int64         `json:"move_by"`
	LastActivity time.Duration `json:"last_activity"`
}

// Returns an array of daily chess games where it is a players turn to act.
func (c *Client) PlayerGamesToMove(username string) ([]PlayerGamesToMove, error) {
	var res map[string][]PlayerGamesToMove

	endpoint := fmt.Sprintf("/player/%v/games/to-move", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return nil, err
	}

	return res["games"], nil
}

// Returns standard multi-game format PGN containing all games for a month for a player.
func (c *Client) PlayerGameArchivePGN(username, year, month string) (string, error) {
	var res string

	endpoint := fmt.Sprintf("/player/%v/games/%v/%v/pgn", username, year, month)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type Player struct {
	Rating   int64  `json:"url"`
	Result   string `json:"result"`
	ID       string `json:"@id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

type PlayerGame struct {
	URL          string        `json:"url"`
	PGN          string        `json:"pgn"`
	TimeControl  string        `json:"time_control"`
	EndTime      time.Duration `json:"end_time"`
	Rated        bool          `json:"rated"`
	TCN          string        `json:"tcn"`
	UUID         string        `json:"uuid"`
	InitialSetup string        `json:"initial_setup"`
	FEN          string        `json:"fen"`
	StartTime    time.Duration `json:"start_time"`
	TimeClass    string        `json:"time_class"`
	Rules        string        `json:"rules"`
	White        Player        `json:"white"`
	Black        Player        `json:"black"`
}

// Returns an array of live and daily chess games that a player has finished.
func (c *Client) PlayerGameMonthlyArchive(username, year, month string) ([]PlayerGame, error) {
	var res map[string][]PlayerGame

	endpoint := fmt.Sprintf("/player/%v/games/%v/%v", username, year, month)

	err := c.get(endpoint, &res)
	if err != nil {
		return nil, err
	}

	return res["games"], nil
}

type PlayerClub struct {
	ID           string        `json:"@id"`
	Name         string        `json:"name"`
	LastActivity time.Duration `json:"last_activity"`
	Icon         string        `json:"icon"`
	URL          string        `json:"url"`
	Joined       time.Duration `json:"joined"`
}

// Returns an array of clubs a player is a member of, with joined date and last activity date.
func (c *Client) PlayerClubs(username string) ([]PlayerClub, error) {
	var res map[string][]PlayerClub

	endpoint := fmt.Sprintf("/player/%v/clubs", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return nil, err
	}

	return res["clubs"], nil
}

type PlayerTournament struct {
	URL           string `json:"url"`
	ID            string `json:"@id"`
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
	Draws         int    `json:"draws"`
	PointsAwarded int    `json:"points_awarded"`
	Placement     int    `json:"placement"`
	Status        string `json:"status"`
	TotalPlayers  int    `json:"total_players"`
}

type PlayerTournamentsResponse struct {
	Finished   []PlayerTournament `json:"finished"`
	InProgress []PlayerTournament `json:"in_progress"`
	Registered []PlayerTournament `json:"registered"`
}

// Returns an array of tournaments a player is registered, is attending or has attended in the past.
func (c *Client) PlayerTournaments(username string) (PlayerTournamentsResponse, error) {
	var res PlayerTournamentsResponse

	endpoint := fmt.Sprintf("/player/%v/tournaments", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Returns an array of titled-player usernames.
func (c *Client) TitledPlayers() ([]string, error) {
	var res map[string][]string

	err := c.get("/titled/GM", &res)
	if err != nil {
		return nil, err
	}

	return res["players"], nil
}
