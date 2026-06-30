package chat

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
)

type ChatImplementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewChatImplementation(chatService service.ChatService) *ChatImplementation {
	return &ChatImplementation{
		chatService: chatService,
	}
}
