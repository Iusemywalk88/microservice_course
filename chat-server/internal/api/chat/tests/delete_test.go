package tests

import (
	"context"
	"fmt"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/api/chat"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
	serviceMock "github.com/Iusemywalk88/microservice_course/chat-server/internal/service/mocks"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"testing"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeleteRequest{
			Id: id,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{{
		name: "success",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: &emptypb.Empty{},
		err:  nil,
		chatServiceMock: func(mc *minimock.Controller) service.ChatService {
			mock := serviceMock.NewChatServiceMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(nil)
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
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatImplementation(chatServiceMock)

			newID, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
