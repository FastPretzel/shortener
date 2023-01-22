package usecase

import (
	"context"
	"errors"
	"shortener/grpc_domain"
	"shortener/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ShortenerUseCase) Create(ctx context.Context, req *grpc_domain.CreateLinkRequest) (*grpc_domain.CreateLinkResponse, error) {
	origLink := req.Link
	if origLink == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	var shortLink string
	for {
		shortLink = hashing(origLink)
		err := s.repo.Create(ctx, origLink, shortLink)
		if err == nil {
			break
		} else if errors.Is(err, domain.ErrUnavailable) {
			return nil, status.Error(codes.Unavailable, domain.ErrUnavailable.Error())
		} else if errors.Is(err, domain.ErrAlreadyExist) {
			existLink, err := s.repo.Get(ctx, shortLink)
			if err != nil {
				if errors.Is(err, domain.ErrUnavailable) {
					return nil, status.Error(codes.Unavailable, domain.ErrUnavailable.Error())
				}
				return nil, err
			}
			if existLink == origLink {
				break
			}
		}
		return nil, err
	}
	return &grpc_domain.CreateLinkResponse{ShortLink: shortLink}, nil
}
