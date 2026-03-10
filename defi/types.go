package defi

import "net/url"

type EventsRequest struct {
	Contract string
	Chain    string
	PageSize int
	Cursor   string
}

func (r EventsRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("contract", r.Contract)
	q.Set("chain", r.Chain)
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type DefiEvent struct {
	TxHash    string         `json:"txHash"`
	Contract  string         `json:"contract"`
	EventType string         `json:"eventType"`
	Timestamp int64          `json:"timestamp"`
	Data      map[string]any `json:"data"`
}

type BlocksRequest struct {
	Chain    string
	PageSize int
	Cursor   string
}

func (r BlocksRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type DefiBlock struct {
	Number    int64  `json:"number"`
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
}

type LatestBlockRequest struct {
	Chain string
}

func (r LatestBlockRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type BlockByTimestampRequest struct {
	Chain     string
	Timestamp int64
}

func (r BlockByTimestampRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	q.Set("timestamp", itoa64(r.Timestamp))
	return q
}
