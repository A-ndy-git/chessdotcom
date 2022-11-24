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
	defaultBaseUrl = "https://api.chess.com/pub"
)

type Client struct {
	// The net/http client
	client *http.Client

	// The API base URL.
	baseUrl string

	// The maximum timeout for any request made.
	timeout time.Duration
}

// Creates a new instance of the ChessDotCom client.
func New() *Client {

	return &Client{
		client:  &http.Client{},
		baseUrl: defaultBaseUrl,
		timeout: defaultTimeout,
	}
}

// Sends a GET request to a specified path. As GET is the only method used, there is only one method handler.
func (c *Client) get(path string, resp interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseUrl, path)

	req, err := c.client.Get(endpoint)
	if err != nil {
		return ErrApiUnavailable
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrGeneric
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return ErrGeneric
	}

	return nil
}
