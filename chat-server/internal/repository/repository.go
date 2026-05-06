package repository

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/model"
)

type ChatRepository interface {
	Create(ctx context.Context, req *model.Chat) (int64, error)
}
