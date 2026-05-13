package chat

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
