package chessdotcom

import (
	"fmt"
	"time"
)

type PlayerProfileResponse struct {
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

func (c *Client) PlayerProfile(username string) (PlayerProfileResponse, error) {
	var res PlayerProfileResponse

	endpoint := fmt.Sprintf("/player/%v", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type LastResponse struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
	Rd     int64         `json:"rd"`
}

type BestResponse struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
	Game   string        `json:"game"`
}

type RecordResponse struct {
	Win            int64         `json:"win"`
	Loss           int64         `json:"loss"`
	Draw           int64         `json:"draw"`
	TimePerMove    time.Duration `json:"time_per_move"`
	TimeoutPercent int           `json:"timeout_percent"`
}

type TournamentResponse struct {
	Points        int64 `json:"points"`
	Withdraw      int64 `json:"withdraw"`
	Count         int64 `json:"count"`
	HighestFinish int64 `json:"highest_finish"`
}

type ChessDailyResponse struct {
	Last       LastResponse       `json:"last"`
	Best       BestResponse       `json:"best"`
	Record     RecordResponse     `json:"record"`
	Tournament TournamentResponse `json:"tournament"`
}

type ChessRapidResponse struct {
	Last   LastResponse   `json:"last"`
	Best   BestResponse   `json:"best"`
	Record RecordResponse `json:"record"`
}

type ChessRating struct {
	Rating int64         `json:"rating"`
	Date   time.Duration `json:"date"`
}

type TacticsResponse struct {
	Highest ChessRating `json:"highest"`
	Lowest  ChessRating `json:"lowest"`
}

type PuzzleRushBestResponse struct {
	TotalAttempts int64 `json:"total_attempts"`
	Score         int64 `json:"score"`
}

type PuzzleRushResponse struct {
	Best PuzzleRushBestResponse `json:"best"`
}

type PlayerStatsResponse struct {
	ChessDaily    ChessDailyResponse     `json:"chess_daily"`
	Chess960Daily ChessDailyResponse     `json:"chess960_daily"`
	ChessRapid    ChessRapidResponse     `json:"chess_rapid"`
	ChessBullet   ChessRapidResponse     `json:"chess_bullet"`
	ChessBlitz    ChessRapidResponse     `json:"chess_blitz"`
	Fide          int64                  `json:"fide"`
	Tactics       TacticsResponse        `json:"tactics"`
	PuzzleRush    PuzzleRushBestResponse `json:"puzzle_rush"`
}

func (c *Client) PlayerStats(username string) (PlayerStatsResponse, error) {
	var res PlayerStatsResponse

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

type PlayerGamesResponse struct {
	Games []PlayerGames `json:"games"`
}

func (c *Client) PlayerGames(username string) (PlayerGamesResponse, error) {
	var res PlayerGamesResponse

	endpoint := fmt.Sprintf("/player/%v/games", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type PlayerGameArchivesResponse struct {
	Archives []string `json:"archives"`
}

func (c *Client) PlayerGameArchives(username string) (PlayerGameArchivesResponse, error) {
	var res PlayerGameArchivesResponse

	endpoint := fmt.Sprintf("/player/%v/games/archives", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type PlayerGamesToMove struct {
	URL          string        `json:"url"`
	MoveBy       int64         `json:"move_by"`
	LastActivity time.Duration `json:"last_activity"`
}

type PlayerGamesToMoveResponse struct {
	Games []PlayerGamesToMove `json:"games"`
}

func (c *Client) PlayerGamesToMove(username string) (PlayerGamesToMoveResponse, error) {
	var res PlayerGamesToMoveResponse

	endpoint := fmt.Sprintf("/player/%v/games/to-move", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

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

type PlayerGameMonthlyArchiveResponse struct {
	Games []PlayerGame `json:"games"`
}

func (c *Client) PlayerGameMonthlyArchive(username, year, month string) (PlayerGameMonthlyArchiveResponse, error) {
	var res PlayerGameMonthlyArchiveResponse

	endpoint := fmt.Sprintf("/player/%v/games/%v/%v", username, year, month)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type PlayerClub struct {
	ID           string        `json:"@id"`
	Name         string        `json:"name"`
	LastActivity time.Duration `json:"last_activity"`
	Icon         string        `json:"icon"`
	URL          string        `json:"url"`
	Joined       time.Duration `json:"joined"`
}

type PlayerClubsResponse struct {
	Clubs []PlayerClub `json:"clubs"`
}

func (c *Client) PlayerClubs(username string) (PlayerClubsResponse, error) {
	var res PlayerClubsResponse

	endpoint := fmt.Sprintf("/player/%v/clubs", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
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

func (c *Client) PlayerTournaments(username string) (PlayerTournamentsResponse, error) {
	var res PlayerTournamentsResponse

	endpoint := fmt.Sprintf("/player/%v/tournaments", username)

	err := c.get(endpoint, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

type TitledPlayersResponse struct {
	Players []string `json:"players"`
}

func (c *Client) TitledPlayers() (TitledPlayersResponse, error) {
	var res TitledPlayersResponse

	err := c.get("/titled/GM", &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
