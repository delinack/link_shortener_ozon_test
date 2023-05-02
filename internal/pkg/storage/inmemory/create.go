package inmemory

import (
	"context"
	"time"

	"link_shorter/internal/pkg/model"

	"github.com/google/uuid"
)

func (i *inMemory) fillURLMap(url, token string) *model.Link {
	i.urlMutex.Lock()
	defer i.urlMutex.Unlock()

	v := &model.Link{
		ID:        uuid.New(),
		URL:       url,
		Token:     token,
		CreatedAt: time.Now(),
	}
	i.urlMap[url] = v

	return v
}

func (i *inMemory) Create(ctx context.Context, url, token string) (*model.Link, error) {
	v := i.fillURLMap(url, token)

	i.tokenMutex.Lock()
	i.tokenMap[token] = v
	i.tokenMutex.Unlock()

	return v, nil
}
