package nameservice

import "net/url"

type ResolveRequest struct {
	Name  string
	Chain string
}

func (r ResolveRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("name", r.Name)
	if r.Chain != "" {
		q.Set("chain", r.Chain)
	}
	return q
}

type ResolveResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
