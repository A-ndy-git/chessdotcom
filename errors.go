package chessdotcom

import "errors"

var (
	ErrApiUnavailable = errors.New("api currently unavailable")
	ErrRateLimit      = errors.New("rate limited by api")
	ErrTimeout        = errors.New("request timed out")
	ErrGeneric        = errors.New("internal server error")
)