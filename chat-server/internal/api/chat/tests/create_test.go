package tests

import (
	"context"
	"fmt"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/api/chat"
	"strconv"

	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"

	serviceMock "github.com/Iusemywalk88/microservice_course/chat-server/internal/service/mocks"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		userID    = gofakeit.Int64()
		usernames = []string{strconv.FormatInt(userID, 10)}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Usernames: usernames,
		}

		chatModel = &model.Chat{
			UserIDs: []int64{userID},
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
		chatServiceMock chatServiceMockFunc
	}{{
		name: "success",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: res,
		err:  nil,
		chatServiceMock: func(mc *minimock.Controller) service.ChatService {
			mock := serviceMock.NewChatServiceMock(mc)
			mock.CreateMock.Expect(ctx, chatModel).Return(id, nil)
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
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatModel).Return(0, serviceErr)
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

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
