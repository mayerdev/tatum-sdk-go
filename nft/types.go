package nft

import "net/url"

type CollectionTokensRequest struct {
	Collection string
	Chain      string
	PageSize   int
	Cursor     string
}

func (r CollectionTokensRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("collection", r.Collection)
	q.Set("chain", r.Chain)
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type NFTToken struct {
	TokenID  string         `json:"tokenId"`
	Contract string         `json:"contract"`
	Chain    string         `json:"chain"`
	Metadata map[string]any `json:"metadata"`
	Owner    string         `json:"owner"`
}

type MetadataRequest struct {
	TokenID  string
	Contract string
	Chain    string
}

func (r MetadataRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("tokenId", r.TokenID)
	q.Set("contract", r.Contract)
	q.Set("chain", r.Chain)
	return q
}

type OwnersRequest struct {
	TokenID  string
	Contract string
	Chain    string
}

func (r OwnersRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("tokenId", r.TokenID)
	q.Set("contract", r.Contract)
	q.Set("chain", r.Chain)
	return q
}

type OwnersResponse struct {
	Owners []string `json:"owners"`
}

type WalletBalancesRequest struct {
	Address  string
	Chain    string
	PageSize int
	Cursor   string
}

func (r WalletBalancesRequest) toQuery() url.Values {
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
