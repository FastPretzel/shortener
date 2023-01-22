package domain

import "errors"

var ErrNotFound = errors.New("link not found")
var ErrUnavailable = errors.New("unavailable")
var ErrAlreadyExist = errors.New("already exist")
