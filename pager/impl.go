package pager

import (
	"net/url"
	"strconv"
)

type Paginated struct {
	PageSize uint64 // 1..50
	Offset   uint64
}

func (p Paginated) AssignTo(query *url.Values) {
	if p.PageSize > 0 {
		query.Set("pageSize", strconv.FormatUint(p.PageSize, 10))
	}

	if p.Offset > 0 {
		query.Set("offset", strconv.FormatUint(p.Offset, 10))
	}
}

func FetchAll[T any](p *Paginated, fetch func() ([]T, error)) ([]T, error) {
	p.PageSize = 50
	p.Offset = 0

	var err error
	var part []T

	result := make([]T, 0)
	for {
		part, err = fetch()
		if err != nil {
			return nil, err
		}

		result = append(result, part...)
		p.Offset += p.PageSize

		if uint64(len(part)) < p.PageSize {
			break
		}
	}

	return result, nil
}
