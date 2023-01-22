package repo

import (
	"context"
	"shortener/internal/domain"
)

func (im *repoIM) Get(_ context.Context, shortLink string) (string, error) {
	v, ok := im.mem[shortLink]
	if ok == false {
		return "", domain.ErrNotFound
	}
	return v, nil
}

func (im *repoIM) Create(_ context.Context, origLink, shortLink string) error {
	im.mu.Lock()
	defer im.mu.Unlock()

	if _, ok := im.mem[shortLink]; ok == true {
		return domain.ErrAlreadyExist
	}

	im.mem[shortLink] = origLink
	return nil
}
