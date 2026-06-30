package redis

import (
	"context"
	"fmt"

	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/cache"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat/redis/converter"

	"time"
)

type repo struct {
	cl         cache.RedisClient
	expiration time.Duration
}

func NewRepository(cl cache.RedisClient, expiration time.Duration) repository.ChatRepository {
	return &repo{
		cl:         cl,
		expiration: expiration,
	}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	chatRepo := converter.ToRepoFromChat(chat)
	idStr := fmt.Sprintf("chat:%d", chat.ID)

	err := r.cl.HashSet(ctx, idStr, chatRepo)
	if err != nil {
		return 0, err
	}

	err = r.cl.Expire(ctx, idStr, r.expiration)
	if err != nil {
		return 0, err
	}
	return chat.ID, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	idStr := fmt.Sprintf("chat:%d", id)
	err := r.cl.Delete(ctx, idStr)
	if err != nil {
		return err
	}
	return nil
}
