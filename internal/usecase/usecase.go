package usecase

import (
	"shortener/grpc_domain"
	"shortener/internal/repo"
)

type ShortenerUseCase struct {
	grpc_domain.UnimplementedShortenerServer
	repo repo.Repository
}

func New(r repo.Repository) *ShortenerUseCase {
	return &ShortenerUseCase{repo: r}
}
