package inmemory

import (
	"context"
	"sync"

	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
)

type inMemory struct {
	urlMutex   sync.Mutex
	tokenMutex sync.Mutex

	urlMap   map[string]*model.Link // create
	tokenMap map[string]*model.Link // get
}

func (i *inMemory) Shutdown(ctx context.Context) {}

func NewInMemory() storage.Store {
	return &inMemory{
		urlMap:   make(map[string]*model.Link),
		tokenMap: make(map[string]*model.Link),
	}
}
