package storage

import (
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
)

type Store interface {
	Create(ctx context.Context, url, token string) (*model.Link, error)
	GetByURL(ctx context.Context, url string) (*model.Link, error)
	GetByToken(ctx context.Context, token string) (*model.Link, error)
	Shutdown(ctx context.Context)
}
