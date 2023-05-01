package service

import (
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
)

type linkService struct {
	store storage.Store
}

type LinkService interface {
	CreateLink(ctx context.Context, url string) (*model.Link, error)
	GetLink(ctx context.Context, token string) (*model.Link, error)
}

func NewLinkService(store storage.Store) LinkService {
	return &linkService{store: store}
}
