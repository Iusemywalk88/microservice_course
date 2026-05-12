package chat

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
)

func (s serv) Create(ctx context.Context, user *model.Chat) (int64, error) {
	id, err := s.chatRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
