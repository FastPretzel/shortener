package repo

import (
	"context"
	"errors"
)

func (im *repoIM) Get(_ context.Context, shortLink string) (string, error) {
	v, ok := im.mem[shortLink]
	if ok == false {
		return "", errors.New("not found link")
	}
	return v, nil
}

func (im *repoIM) Create(_ context.Context, origLink, shortLink string) error {
	im.mu.Lock()
	defer im.mu.Unlock()

	if _, ok := im.mem[shortLink]; ok == true {
		return errors.New("already exists")
	}

	im.mem[shortLink] = origLink
	return nil
}
