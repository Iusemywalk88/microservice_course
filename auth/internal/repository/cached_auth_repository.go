package repository

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
)

type cachedAuthRepository struct {
	pgRepo    AuthRepository
	redisRepo AuthRepository
}

func NewCachedAuthRepository(pgRepo AuthRepository, redisRepo AuthRepository) AuthRepository {
	return &cachedAuthRepository{
		pgRepo:    pgRepo,
		redisRepo: redisRepo,
	}
}

func (c *cachedAuthRepository) Create(ctx context.Context, req *model.User) (int64, error) {
	id, err := c.pgRepo.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	req.ID = id
	_, _ = c.redisRepo.Create(ctx, req)

	return id, nil
}

func (c *cachedAuthRepository) Delete(ctx context.Context, id int64) error {
	err := c.pgRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_ = c.redisRepo.Delete(ctx, id)

	return nil
}

func (c *cachedAuthRepository) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := c.redisRepo.Get(ctx, id)
	if err == nil && user != nil {
		return user, nil
	}

	user, err = c.pgRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if user != nil {
		_, _ = c.redisRepo.Create(ctx, user)
	}

	return user, nil
}

func (c *cachedAuthRepository) Update(ctx context.Context, user *model.UpdateUserInfo) error {
	err := c.pgRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	_ = c.redisRepo.Delete(ctx, user.ID)

	return nil
}
