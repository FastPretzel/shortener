package usecase

import (
	"context"
	"shortener/grpc_domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ShortenerUseCase) Get(ctx context.Context, req *grpc_domain.GetLinkRequest) (*grpc_domain.GetLinkResponse, error) {
	shortLink := req.Link
	if shortLink == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	return &grpc_domain.GetLinkResponse{OrigLink: "Hello"}, nil
}
