package repo

import (
	"context"
	"shortener/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	testTable := []struct {
		name         string
		arg          string
		expectedResp string
		expectedErr  error
		notFound     bool
	}{
		{
			name:         "OK",
			arg:          "aBcDeF123_",
			expectedResp: "test.com",
			expectedErr:  nil,
			notFound:     false,
		},
		{
			name:         "Expect error not found",
			arg:          "aBcDeF123_",
			expectedResp: "",
			expectedErr:  domain.ErrNotFound,
			notFound:     true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			im := NewIM()
			if test.notFound == false {
				im.Create(context.Background(), test.expectedResp, test.arg)
			}
			resp, err := im.Get(context.Background(), test.arg)

			assert.Equal(t, test.expectedResp, resp)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestCreate(t *testing.T) {
	testTable := []struct {
		name         string
		origLink     string
		shortLink    string
		expectedErr  error
		alreadyExist bool
	}{
		{
			name:         "OK",
			origLink:     "test.com",
			shortLink:    "aBcDeF123_",
			expectedErr:  nil,
			alreadyExist: false,
		},
		{
			name:         "Expect error already exist",
			origLink:     "test.com",
			shortLink:    "aBcDeF123_",
			expectedErr:  domain.ErrAlreadyExist,
			alreadyExist: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			im := NewIM()
			if test.alreadyExist == true {
				im.Create(context.Background(), test.origLink, test.shortLink)
			}
			err := im.Create(context.Background(),
				test.origLink, test.shortLink)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
