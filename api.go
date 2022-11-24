package chessdotcom

import (
	"fmt"
	"time"
)

/*
 * Player Endpoints.
 */

type PlayerProfileResponse struct {
	Avatar     string        `json:"avatar"`
	PlayerId   int64         `json:"player_id"`
	Id         string        `json:"id"`
	Url        string        `json:"url"`
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

	err := c.get(fmt.Sprintf("/player/%v", username), &res)
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

	err := c.get(fmt.Sprintf("/player/%v/stats", username), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
