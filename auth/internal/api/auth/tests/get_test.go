package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Iusemywalk88/microservice_course/auth/internal/api/auth"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	serviceMock "github.com/Iusemywalk88/microservice_course/auth/internal/service/mocks"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		userRole  = model.Role(gofakeit.IntRange(0, 1))
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		user = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			UserRole:  userRole,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{updatedAt, true},
		}

		res = &desc.GetResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      desc.Role(userRole),
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		authServiceMock authServiceMockFunc
	}{{
		name: "success",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: res,
		err:  nil,
		authServiceMock: func(mc *minimock.Controller) service.AuthService {
			mock := serviceMock.NewAuthServiceMock(mc)
			mock.GetMock.Expect(ctx, id).Return(user, nil)
			return mock
		},
	},
		{
			name: "auth service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewAuthImplementation(authServiceMock)

			newID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)

		})
	}
}
