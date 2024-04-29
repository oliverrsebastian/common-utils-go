package database

import (
	"context"
	"gorm.io/gorm"
	"math"
)

type Paging struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalPage int   `json:"total_page"`
	TotalSize int64 `json:"total_size"`
}

type PageRequest struct {
	Page   int
	Size   int
	SortBy string
}

func PaginateByOffset(ctx context.Context, tx *gorm.DB, req *PageRequest, dest interface{}, sortMap ...SortMap) (*Paging, error) {
	var count int64
	errCh := make(chan error, 1)

	st := tx.Statement
	preloads := st.Preloads
	st.Preloads = nil
	tx.Statement = st

	go countRecords(tx, &count, errCh)
	if err := <-errCh; err != nil {
		return nil, err
	}

	st.Preloads = preloads
	tx.Statement = st

	page := req.Page
	size := req.Size
	offset := page * size

	query := tx.Offset(offset).Limit(size)
	query = applySort(tx, req, sortMap...)

	if err := query.Find(dest).Error; err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(count) / float64(size)))
	return &Paging{
		Page:      page,
		Size:      size,
		TotalSize: count,
		TotalPage: totalPage,
	}, nil
}

func countRecords(tx *gorm.DB, count *int64, errCh chan error) {
	errCh <- tx.Count(count).Error
}
