package inmemory

import (
	"context"

	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
)

func (i *inMemory) GetByToken(ctx context.Context, token string) (*model.Link, error) {
	i.tokenMutex.Lock()
	defer i.tokenMutex.Unlock()
	v, ok := i.tokenMap[token]
	if !ok {
		return nil, storage.ErrNotFound
	}
	return v, nil

}

func (i *inMemory) GetByURL(ctx context.Context, url string) (*model.Link, error) {
	i.urlMutex.Lock()
	defer i.urlMutex.Unlock()
	v, ok := i.urlMap[url]
	if !ok {
		return nil, storage.ErrNotFound
	}
	return v, nil

}
