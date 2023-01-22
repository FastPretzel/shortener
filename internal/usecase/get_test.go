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

func TestGet(t *testing.T) {
	type mockBehavior func(r *mock_repo.MockRepository, shortLink string)

	testTable := []struct {
		name         string
		arg          *grpc_domain.GetLinkRequest
		mockBehavior mockBehavior
		expectedResp *grpc_domain.GetLinkResponse
		expectedErr  error
	}{
		{
			name: "OK",
			arg:  &grpc_domain.GetLinkRequest{Link: "aBcDeF123_"},
			mockBehavior: func(r *mock_repo.MockRepository, shortLink string) {
				r.EXPECT().Get(context.Background(), shortLink).Return("test.com", nil)
			},
			expectedResp: &grpc_domain.GetLinkResponse{OrigLink: "test.com"},
			expectedErr:  nil,
		},
		{
			name:         "Expect error bad request",
			arg:          &grpc_domain.GetLinkRequest{Link: "aBc"},
			mockBehavior: func(r *mock_repo.MockRepository, shortLink string) {},
			expectedResp: nil,
			expectedErr:  status.Error(codes.InvalidArgument, "Bad request"),
		},
		{
			name: "Expect error not found",
			arg:  &grpc_domain.GetLinkRequest{Link: "aBcDeF123_"},
			mockBehavior: func(r *mock_repo.MockRepository, shortLink string) {
				r.EXPECT().Get(context.Background(), shortLink).Return("", domain.ErrNotFound)
			},
			expectedResp: nil,
			expectedErr:  status.Error(codes.NotFound, domain.ErrNotFound.Error()),
		},
		{
			name: "Expect error unavailable",
			arg:  &grpc_domain.GetLinkRequest{Link: "aBcDeF123_"},
			mockBehavior: func(r *mock_repo.MockRepository, shortLink string) {
				r.EXPECT().Get(context.Background(), shortLink).Return("", domain.ErrUnavailable)
			},
			expectedResp: nil,
			expectedErr:  status.Error(codes.Unavailable, domain.ErrUnavailable.Error()),
		},
		{
			name: "Expect another error",
			arg:  &grpc_domain.GetLinkRequest{Link: "aBcDeF123_"},
			mockBehavior: func(r *mock_repo.MockRepository, shortLink string) {
				r.EXPECT().Get(context.Background(), shortLink).Return("", errors.New("error"))
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
			test.mockBehavior(repo, test.arg.Link)

			uc := New(repo)

			resp, err := uc.Get(context.Background(), test.arg)

			assert.Equal(t, test.expectedResp, resp)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
