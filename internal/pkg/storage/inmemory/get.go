package inmemory

import (
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
)

func (i *inMemory) GetByToken(ctx context.Context, token string) (*model.Link, error) {
	i.tokenMutex.Lock()
	defer i.tokenMutex.Unlock()
	if v, ok := i.tokenMap[token]; !ok {
		return nil, storage.ErrNotFound
	} else {
		return v, nil
	}
}

func (i *inMemory) GetByURL(ctx context.Context, url string) (*model.Link, error) {
	i.urlMutex.Lock()
	defer i.urlMutex.Unlock()
	if v, ok := i.urlMap[url]; !ok {
		return nil, storage.ErrNotFound
	} else {
		return v, nil
	}
}
