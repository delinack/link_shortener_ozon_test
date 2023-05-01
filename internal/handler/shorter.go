package handler

import (
	"link_shorter/internal/pkg/service"
	"link_shorter/internal/protobuf/link_shorter/protobuf/shorter"
)

type Implementation struct {
	shorter.UnimplementedShorterServiceServer

	linkService service.LinkService
}

func NewShorter(linkService service.LinkService) *Implementation {
	return &Implementation{linkService: linkService}
}
