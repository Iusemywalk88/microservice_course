package service

import (
	"context"

	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, req *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
