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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		userRole = model.Role(gofakeit.IntRange(0, 1))

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  wrapperspb.String(name),
			Email: wrapperspb.String(email),
			Role:  desc.Role(userRole),
		}

		user = &model.UpdateUserInfo{
			ID:       id,
			Name:     &name,
			Email:    &email,
			UserRole: userRole,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		authServiceMock authServiceMockFunc
	}{{
		name: "success",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: &emptypb.Empty{},
		err:  nil,
		authServiceMock: func(mc *minimock.Controller) service.AuthService {
			mock := serviceMock.NewAuthServiceMock(mc)
			mock.UpdateMock.Expect(ctx, user).Return(nil)
			return mock
		},
	},
		{
			name: "auth service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, user).Return(serviceErr)
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

			newID, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
