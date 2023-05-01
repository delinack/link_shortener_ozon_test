package service

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"link_shorter/internal/pkg/model"
	"link_shorter/internal/pkg/storage"
	"strings"
)

const (
	alphabet   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	hashLength = 10
)

func tokenGenerate() string {
	id := uuid.New()
	idBytes := []byte(id.String())
	hash := md5.Sum(idBytes)

	var indices [hashLength]int
	for i := 0; i < hashLength; i++ {
		index := int(hash[i]) % len(alphabet)
		indices[i] = index
	}

	token := make([]byte, hashLength)
	for i, index := range indices {
		token[i] = alphabet[index]
	}

	return string(token)
}

func (s *linkService) CreateLink(ctx context.Context, url string) (*model.Link, error) {
	url = strings.Trim(url, " ")
	link, err := s.store.GetByURL(ctx, url)
	if err != nil && err != storage.ErrNotFound {
		return nil, fmt.Errorf("store.GetByURL: %w", err)
	}
	if link != nil {
		return link, nil
	}

	token := tokenGenerate()
	v, err := s.store.Create(ctx, url, token)
	if err != nil {
		return nil, fmt.Errorf("store.Create: %w", err)
	}

	return v, nil
}
