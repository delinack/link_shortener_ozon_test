package service

import (
	"fmt"
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
)

func (s *linkService) GetLink(ctx context.Context, token string) (*model.Link, error) {
	v, err := s.store.GetByToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("store.Get: %w", err)
	}

	return v, nil
}
