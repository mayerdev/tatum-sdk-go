package tatum

import (
	"net/http"
	"time"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type clientConfig struct {
	baseURL    string
	rpcBaseURL string
	httpClient httpclient.Doer
	maxRetries int
	userAgent  string
}

type ClientOption func(*clientConfig)

func WithBaseURL(u string) ClientOption {
	return func(c *clientConfig) { c.baseURL = u }
}

func WithRPCBaseURL(u string) ClientOption {
	return func(c *clientConfig) { c.rpcBaseURL = u }
}

func WithHTTPClient(hc httpclient.Doer) ClientOption {
	return func(c *clientConfig) { c.httpClient = hc }
}

func WithTimeout(d time.Duration) ClientOption {
	return func(c *clientConfig) { c.httpClient = &http.Client{Timeout: d} }
}

func WithRetries(n int) ClientOption {
	return func(c *clientConfig) { c.maxRetries = n }
}

func WithUserAgent(ua string) ClientOption {
	return func(c *clientConfig) { c.userAgent = ua }
}
