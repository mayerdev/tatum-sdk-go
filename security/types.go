package security

import "net/url"

type CheckAddressRequest struct {
	Address string
	Chain   string
}

func (r CheckAddressRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type CheckAddressResponse struct {
	Address   string   `json:"address"`
	Chain     string   `json:"chain"`
	Malicious bool     `json:"malicious"`
	Tags      []string `json:"tags"`
}
