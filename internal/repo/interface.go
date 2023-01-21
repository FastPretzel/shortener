package repo

import "context"

type Repository interface {
	Get(ctx context.Context, shortLink string) (string, error)
	Create(ctx context.Context, origLink, shortLink string) error
}
