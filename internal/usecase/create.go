package usecase

import (
	"context"
	"shortener/grpc_domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ShortenerUseCase) Create(ctx context.Context, req *grpc_domain.CreateLinkRequest) (*grpc_domain.CreateLinkResponse, error) {
	origLink := req.Link
	if origLink == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	return &grpc_domain.CreateLinkResponse{ShortLink: "Hello"}, nil
}
