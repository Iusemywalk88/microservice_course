package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
)

func (i *ChatImplementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("Delete request: %v", req.GetId())

	return &emptypb.Empty{}, nil
}
