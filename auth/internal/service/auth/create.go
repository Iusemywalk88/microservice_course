package auth

import (
	"context"

	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
)

func (s serv) Create(ctx context.Context, req *model.User) (int64, error) {
	id, err := s.authRepo.Create(ctx, req)
	if err != nil {
		return 0, err
	}
	return id, nil
}
