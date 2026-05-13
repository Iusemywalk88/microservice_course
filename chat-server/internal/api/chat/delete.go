package chat

import (
	"context"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("Delete request: %v", req.GetId())

	return &emptypb.Empty{}, nil
}
