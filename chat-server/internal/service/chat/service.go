package chat

import (
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
)

type serv struct {
	chatRepo  repository.ChatRepository
	txManager db.TxManager
}

func NewService(chatRepo repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{chatRepo: chatRepo, txManager: txManager}
}
