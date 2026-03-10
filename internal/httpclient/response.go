package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func decodeResponse(resp *http.Response, out any) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return parseAPIError(resp.StatusCode, data)
	}

	if out == nil {
		return nil
	}

	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func parseAPIError(statusCode int, data []byte) error {
	var payload struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
		Data      any    `json:"data"`
	}
	_ = json.Unmarshal(data, &payload)
	return &APIError{
		StatusCode: statusCode,
		ErrorCode:  payload.ErrorCode,
		Message:    payload.Message,
		Data:       payload.Data,
	}
}

type APIError struct {
	StatusCode int
	ErrorCode  string
	Message    string
	Data       any
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("tatum api error %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("tatum api error %d", e.StatusCode)
}
