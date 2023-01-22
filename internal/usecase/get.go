package usecase

import (
	"context"
	"errors"
	"shortener/grpc_domain"
	"shortener/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ShortenerUseCase) Get(ctx context.Context, req *grpc_domain.GetLinkRequest) (*grpc_domain.GetLinkResponse, error) {
	shortLink := req.Link
	if shortLink == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	origLink, err := s.repo.Get(ctx, shortLink)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, domain.ErrNotFound.Error())
		} else if errors.Is(err, domain.ErrUnavailable) {
			return nil, status.Error(codes.Unavailable, domain.ErrUnavailable.Error())
		}
		return nil, err
	}
	return &grpc_domain.GetLinkResponse{OrigLink: origLink}, nil
}
