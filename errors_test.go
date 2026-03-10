package tatum_test

import (
	"testing"

	tatum "gitlab.com/mayerdev/tatum-sdk-go"
)

func TestIsNotFound(t *testing.T) {
	err := &tatum.APIError{StatusCode: 404, Message: "not found"}
	if !tatum.IsNotFound(err) {
		t.Error("expected IsNotFound to return true")
	}
	if tatum.IsRateLimit(err) {
		t.Error("expected IsRateLimit to return false")
	}
}

func TestIsRateLimit(t *testing.T) {
	err := &tatum.APIError{StatusCode: 429, Message: "rate limit"}
	if !tatum.IsRateLimit(err) {
		t.Error("expected IsRateLimit to return true")
	}
	if tatum.IsNotFound(err) {
		t.Error("expected IsNotFound to return false")
	}
}

func TestAPIError_Error(t *testing.T) {
	err := &tatum.APIError{StatusCode: 500, Message: "internal error"}
	if err.Error() == "" {
		t.Error("expected non-empty error message")
	}
}
