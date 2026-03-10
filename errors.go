package tatum

import (
	"errors"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type APIError = httpclient.APIError

func IsNotFound(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == 404
}

func IsRateLimit(err error) bool {
	var e *APIError
	return errors.As(err, &e) && e.StatusCode == 429
}
