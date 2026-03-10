package transactions

import "net/url"

type HistoryRequest struct {
	Address  string
	Chain    string
	PageSize int
	Cursor   string
}

func (r HistoryRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type Transaction struct {
	Hash        string `json:"hash"`
	BlockNumber int64  `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Fee         string `json:"fee"`
	Chain       string `json:"chain"`
	Timestamp   int64  `json:"timestamp"`
}

type ByHashRequest struct {
	Hash  string
	Chain string
}

func (r ByHashRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("hash", r.Hash)
	q.Set("chain", r.Chain)
	return q
}
