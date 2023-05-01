package handler

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"link_shorter/internal/config"
	"link_shorter/internal/pkg/service"
	"link_shorter/internal/pkg/storage"
	"link_shorter/internal/pkg/storage/inmemory"
	"link_shorter/internal/pkg/storage/postgres"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"
	"testing"
)

func TestUrlValidate(t *testing.T) {
	testCases := []struct {
		link          string
		expectedError error
	}{
		{
			link:          "https://www.delinack.ru/",
			expectedError: nil,
		},
		{
			link:          "asdasd",
			expectedError: fmt.Errorf("%s %q: %s", "parse", "asdasd", "invalid URI for request"),
		},
		{
			link:          "",
			expectedError: fmt.Errorf("empty string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.link, func(t *testing.T) {
			err := urlValidate(tc.link)
			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected Error: %q, got: %q", tc.expectedError, err)
			}
		})
	}
}

func TestCreateLink(t *testing.T) {
	tests := func(store storage.Store) {
		var ctx context.Context
		linkService := service.NewLinkService(store)
		implementation := NewShorter(linkService)

		// создаём токен
		token, err := implementation.CreateLink(ctx, &shorter.CreateLinkRequest{Url: testURL})
		require.NoError(t, err)
		// создаём токен ещё раз по той же ссылке
		token1, err := implementation.CreateLink(ctx, &shorter.CreateLinkRequest{Url: testURL})
		require.NoError(t, err)
		// проверяем токены на совпадение
		require.Equal(t, token, token1)
	}

	tests(inmemory.NewInMemory())
	tests(postgres.NewPostgres(config.InitTests()))
}

// TODO: create lint test
