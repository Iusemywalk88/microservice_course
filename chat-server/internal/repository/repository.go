package repository

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, req *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
