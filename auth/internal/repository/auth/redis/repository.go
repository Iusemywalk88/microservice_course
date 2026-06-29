package redis

import (
	"context"
	"fmt"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/cache"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/redis/converter"
	modelRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/redis/model"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type repo struct {
	cl         cache.RedisClient
	expiration time.Duration
}

func NewRepository(cl cache.RedisClient, expiration time.Duration) repository.AuthRepository {
	return &repo{
		cl:         cl,
		expiration: expiration,
	}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	userAuth := converter.ToRepoFromUser(user)
	idStr := fmt.Sprintf("user:%d", user.ID)

	err := r.cl.HashSet(ctx, idStr, userAuth)
	if err != nil {
		return 0, err
	}

	err = r.cl.Expire(ctx, idStr, r.expiration)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	idStr := fmt.Sprintf("user:%d", id)
	values, err := r.cl.HGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, nil
	}

	var user modelRepo.User

	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, user *model.UpdateUserInfo) error {
	idStr := fmt.Sprintf("user:%d", user.ID)
	err := r.cl.Delete(ctx, idStr)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	idStr := fmt.Sprintf("user:%d", id)
	err := r.cl.Delete(ctx, idStr)
	if err != nil {
		return err
	}
	return nil
}
