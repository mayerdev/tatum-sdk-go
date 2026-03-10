package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient Doer
	MaxRetries int
	UserAgent  string
}

type Client struct {
	cfg Config
}

func New(cfg Config) *Client {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = 3
	}
	if cfg.UserAgent == "" {
		cfg.UserAgent = "tatum-sdk-go/1.0"
	}
	return &Client{cfg: cfg}
}

func (c *Client) Get(ctx context.Context, path string, query url.Values, out any) error {
	u := c.cfg.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	return c.do(ctx, http.MethodGet, u, nil, out)
}

func (c *Client) Post(ctx context.Context, path string, query url.Values, body any, out any) error {
	u := c.cfg.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	return c.do(ctx, http.MethodPost, u, body, out)
}

func (c *Client) Delete(ctx context.Context, path string, query url.Values, out any) error {
	u := c.cfg.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	return c.do(ctx, http.MethodDelete, u, nil, out)
}

func (c *Client) PostAbsolute(ctx context.Context, absoluteURL string, body any, out any) error {
	return c.do(ctx, http.MethodPost, absoluteURL, body, out)
}

func (c *Client) do(ctx context.Context, method, rawURL string, body any, out any) error {
	var lastErr error
	for attempt := 0; attempt <= c.cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(100*(1<<(attempt-1))) * time.Millisecond
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}
		}

		var bodyBytes []byte
		if body != nil {
			var err error
			bodyBytes, err = json.Marshal(body)
			if err != nil {
				return fmt.Errorf("marshal request body: %w", err)
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, rawURL, bytes.NewReader(bodyBytes))
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}

		req.Header.Set("x-api-key", c.cfg.APIKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", c.cfg.UserAgent)

		resp, err := c.cfg.HTTPClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusBadGateway ||
			resp.StatusCode == http.StatusServiceUnavailable ||
			resp.StatusCode == http.StatusGatewayTimeout {
			resp.Body.Close()
			lastErr = fmt.Errorf("status %d", resp.StatusCode)
			continue
		}

		return decodeResponse(resp, out)
	}
	return lastErr
}
