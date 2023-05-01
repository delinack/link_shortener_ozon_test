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

const (
	testURL = "https://www.delinack.ru/"
)

func TestTokenValidate(t *testing.T) {
	testCases := []struct {
		token         string
		expectedError error
	}{
		{
			token:         "fg34aF3nb_",
			expectedError: nil,
		},
		{
			token:         "luM4ac3nQB",
			expectedError: nil,
		},
		{
			token:         "2y3e4I9_np",
			expectedError: nil,
		},
		{
			token:         "2y3e4I9_n!",
			expectedError: fmt.Errorf("invalid token"),
		},
		{
			token:         "123",
			expectedError: fmt.Errorf("invalid token"),
		},
		{
			token:         "",
			expectedError: fmt.Errorf("empty string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.token, func(t *testing.T) {
			err := tokenValidate(tc.token)
			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected Error: %q, got: %q", tc.expectedError, err)
			}
		})
	}
}

func TestGetLink(t *testing.T) {
	tests := func(store storage.Store) {
		var ctx context.Context
		linkService := service.NewLinkService(store)
		implementation := NewShorter(linkService)

		// проверяем несуществующий токен
		_, err := implementation.GetLink(ctx, &shorter.GetLinkRequest{Token: "fg34aF3nb_"})
		require.EqualError(t, err, "rpc error: code = NotFound desc = url not found by token")

		// создаём токен
		token, err := implementation.CreateLink(ctx, &shorter.CreateLinkRequest{Url: testURL})
		require.NoError(t, err)
		// проверяем существующий токен
		res, err := implementation.GetLink(ctx, &shorter.GetLinkRequest{Token: token.Token})
		require.NoError(t, err)
		require.Equal(t, testURL, res.Url)
	}

	tests(inmemory.NewInMemory())
	tests(postgres.NewPostgres(config.InitTests()))
}
