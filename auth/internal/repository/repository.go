package repository

import (
	"context"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
)

type AuthRepository interface {
	Get(ctx context.Context, id int64) (*desc.GetResponse, error)
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
}
