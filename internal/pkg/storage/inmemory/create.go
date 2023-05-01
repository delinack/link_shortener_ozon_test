package inmemory

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
	"time"
)

func (i *inMemory) fillURLMap(url, token string) *model.Link {
	i.urlMutex.Lock()
	defer i.urlMutex.Unlock()

	v := &model.Link{
		Id:        uuid.New(),
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
