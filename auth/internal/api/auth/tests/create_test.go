package tests

import (
	"context"
	"fmt"
	"github.com/Iusemywalk88/microservice_course/auth/internal/api/auth"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	serviceMock "github.com/Iusemywalk88/microservice_course/auth/internal/service/mocks"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(false, false, false, false, false, 12)
		userRole = model.Role(gofakeit.IntRange(0, 1))

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role(userRole),
		}

		user = &model.User{
			Name:     name,
			Email:    email,
			Password: password,
			UserRole: userRole,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
			mock.CreateMock.Expect(ctx, user).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, user).Return(0, serviceErr)
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

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, newID)
			}
		})
	}
}
