// API contains the actual methods of the Chess.com API. E.g Get Player.
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
