package usecase

import (
	"context"
	"errors"
	"shortener/grpc_domain"
	"shortener/internal/domain"
	mock_repo "shortener/internal/repo/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreate(t *testing.T) {
	type mockBehavior func(r *mock_repo.MockRepository, origLink, shortLink string)

	testTable := []struct {
		name         string
		arg          *grpc_domain.CreateLinkRequest
		mockBehavior mockBehavior
		expectedResp *grpc_domain.CreateLinkResponse
		expectedErr  error
	}{
		{
			name: "OK",
			arg:  &grpc_domain.CreateLinkRequest{Link: "google.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(nil)
			},
			expectedResp: &grpc_domain.CreateLinkResponse{ShortLink: "huPVAz7b6R"},
			expectedErr:  nil,
		},
		{
			name:         "Expect error bad request",
			arg:          &grpc_domain.CreateLinkRequest{Link: ""},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {},
			expectedResp: nil,
			expectedErr:  status.Error(codes.InvalidArgument, "Bad request"),
		},
		{
			name: "Expect error unavailable",
			arg:  &grpc_domain.CreateLinkRequest{Link: "test.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(domain.ErrUnavailable)
			},
			expectedResp: nil,
			expectedErr:  status.Error(codes.Unavailable, domain.ErrUnavailable.Error()),
		},
		{
			name: "Expect error already exist",
			arg:  &grpc_domain.CreateLinkRequest{Link: "test.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(domain.ErrAlreadyExist)
				r.EXPECT().Get(context.Background(), shortLink).Return("", domain.ErrUnavailable)
			},
			expectedResp: nil,
			expectedErr:  status.Error(codes.Unavailable, domain.ErrUnavailable.Error()),
		},
		{
			name: "Expect another error already exist",
			arg:  &grpc_domain.CreateLinkRequest{Link: "test.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(domain.ErrAlreadyExist)
				r.EXPECT().Get(context.Background(), shortLink).Return("", errors.New("error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
		{
			name: "Link already exists",
			arg:  &grpc_domain.CreateLinkRequest{Link: "test.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(domain.ErrAlreadyExist)
				r.EXPECT().Get(context.Background(), shortLink).Return("test.com", nil)
			},
			expectedResp: &grpc_domain.CreateLinkResponse{ShortLink: hashing("test.com")},
			expectedErr:  nil,
		},
		{
			name: "Expect error unavailable",
			arg:  &grpc_domain.CreateLinkRequest{Link: "test.com"},
			mockBehavior: func(r *mock_repo.MockRepository, origLink, shortLink string) {
				r.EXPECT().Create(context.Background(), origLink, shortLink).Return(errors.New("error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("error"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repo.NewMockRepository(ctrl)
			test.mockBehavior(repo, test.arg.Link, hashing(test.arg.Link))

			uc := New(repo)

			resp, err := uc.Create(context.Background(), test.arg)

			assert.Equal(t, test.expectedResp, resp)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
