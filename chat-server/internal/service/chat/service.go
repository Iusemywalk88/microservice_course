package chat

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
)

type serv struct {
	chatRepo repository.ChatRepository
}

func NewService(chatRepo repository.ChatRepository) service.ChatService {
	return serv{chatRepo: chatRepo}
}
