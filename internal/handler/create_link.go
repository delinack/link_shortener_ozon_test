package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"
	"net/url"
)

func urlValidate(link string) error {
	if link == "" {
		return fmt.Errorf("empty string")
	}

	parsed, err := url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		return fmt.Errorf("invalid url: %s", link)
	}

	return nil
}

func (i *Implementation) CreateLink(ctx context.Context, request *shorter.CreateLinkRequest) (*shorter.CreateLinkResponse, error) {
	if err := urlValidate(request.Url); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "linkService.CreateLink: %v", err)
	}

	v, err := i.linkService.CreateLink(ctx, request.Url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "linkService.CreateLink: %v", err)

	}

	return &shorter.CreateLinkResponse{
		Token: v.Token,
	}, nil
}
