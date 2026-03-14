package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	parts := []string{
		fmt.Sprintf("tatum error %d", e.StatusCode),
	}

	if e.ErrorCode != "" {
		parts = append(parts, " (", e.ErrorCode, ")")
	}

	if e.Message != "" {
		filteredMessage := strings.Replace(e.Message, " Please see data for additional information.", "", 1)
		parts = append(parts, ": ", filteredMessage)
	}

	if e.Data != nil {
		parts = append(parts, fmt.Sprintf(" Data: %v", e.Data))
	}

	return strings.Join(parts, "")
}
