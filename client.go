package chessdotcom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultTimeout = time.Minute
	baseUrl        = "https://api.chess.com/pub"
	userAgent      = "github.com/a-ndy-git/chessdotcom/1.0.0"
	pgn            = "application/vnd.chess-pgn; charset=utf-8; charset=utf-8"
)

type Client struct {
	*http.Client
}

// Creates a new ChessDotCom client instance with default opts.
func New() *Client {

	return &Client{
		&http.Client{Timeout: defaultTimeout},
	}
}

// Sends a GET request to a specified path.
func (c *Client) get(path string, resp interface{}) error {
	endpoint := fmt.Sprintf("%s%s", baseUrl, path)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("UserAgent", userAgent)

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {

		switch res.StatusCode {
		case 301:
		case 304:
		case 404:
		case 410:
		case 429:
		default:
		}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Return as a string if data is of PGN format.
	if res.Header.Get("content-type") == pgn {
		*resp.(*string) = string(body[:])
		return nil
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return nil
}
