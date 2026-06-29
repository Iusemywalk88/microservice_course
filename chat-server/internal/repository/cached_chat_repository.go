package repository

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/model"
)

type cachedChatRepository struct {
	pgRepo    ChatRepository
	redisRepo ChatRepository
}

func NewCachedChatRepository(pgRepo ChatRepository, redisRepo ChatRepository) ChatRepository {
	return &cachedChatRepository{
		pgRepo:    pgRepo,
		redisRepo: redisRepo,
	}
}

func (c *cachedChatRepository) Create(ctx context.Context, req *model.Chat) (int64, error) {
	id, err := c.pgRepo.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	req.ID = id

	_, _ = c.redisRepo.Create(ctx, req)

	return id, nil
}

func (c *cachedChatRepository) Delete(ctx context.Context, id int64) error {
	err := c.pgRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_ = c.redisRepo.Delete(ctx, id)

	return nil
}
