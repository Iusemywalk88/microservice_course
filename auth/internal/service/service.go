package service

import (
	"context"

	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, req *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user *model.UpdateUserInfo) error
}
