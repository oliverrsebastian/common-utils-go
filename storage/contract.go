package storage

import (
	"context"
	"io"
)

type Client interface {
	Upload(ctx context.Context, folder, file string, r io.Reader) (url string, err error)
	Close() error
}
