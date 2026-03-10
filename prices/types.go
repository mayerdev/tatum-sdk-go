package prices

import "net/url"

type RateRequest struct {
	Currency string
	BasePair string
}

func (r RateRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("currency", r.Currency)
	if r.BasePair != "" {
		q.Set("basePair", r.BasePair)
	}
	return q
}

type RateResponse struct {
	Currency  string  `json:"currency"`
	Value     float64 `json:"value"`
	BasePair  string  `json:"basePair"`
	Timestamp int64   `json:"timestamp"`
}

type OHLCVRequest struct {
	Currency string
	BasePair string
	From     int64
	To       int64
}

func (r OHLCVRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("currency", r.Currency)
	if r.BasePair != "" {
		q.Set("basePair", r.BasePair)
	}
	if r.From > 0 {
		q.Set("from", itoa64(r.From))
	}
	if r.To > 0 {
		q.Set("to", itoa64(r.To))
	}
	return q
}

type OHLCVCandle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type PriceChangeRequest struct {
	Currency string
	BasePair string
}

func (r PriceChangeRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("currency", r.Currency)
	if r.BasePair != "" {
		q.Set("basePair", r.BasePair)
	}
	return q
}

type PriceChangeResponse struct {
	Currency      string  `json:"currency"`
	Change24h     float64 `json:"change24h"`
	ChangePercent float64 `json:"changePercent"`
}
