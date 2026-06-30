package chat

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/converter"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
)

func (i *ChatImplementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatModel, err := converter.ToServiceFromDesc(req)
	if err != nil {
		return nil, err
	}

	id, err := i.chatService.Create(ctx, chatModel)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
