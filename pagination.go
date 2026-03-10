package tatum

type Page[T any] struct {
	Data       []T    `json:"data"`
	NextCursor string `json:"nextCursor"`
	PrevCursor string `json:"prevCursor"`
	Total      int64  `json:"total"`
}

func (p Page[T]) HasMore() bool {
	return p.NextCursor != ""
}

type PageRequest struct {
	PageSize int
	Cursor   string
}
