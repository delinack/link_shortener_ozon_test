package handler

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"link_shorter/internal/pkg/storage"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func tokenValidate(token string) error {
	if token == "" {
		return fmt.Errorf("empty string")
	}

	ok, err := regexp.Match("[0-9a-zA-Z_]{10}", []byte(token))
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (i *Implementation) GetLink(ctx context.Context, request *shorter.GetLinkRequest) (*shorter.GetLinkResponse, error) {
	if err := tokenValidate(request.Token); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "linkService.GetLink: %v", err)
	}

	v, err := i.linkService.GetLink(ctx, request.Token)
	if err != nil && errors.Unwrap(err) == storage.ErrNotFound {
		return nil, status.Error(codes.NotFound, "url not found by token")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "linkService.GetLink: %v", err)
	}

	return &shorter.GetLinkResponse{
		Url: v.URL,
	}, nil
}
