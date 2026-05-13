package auth

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
)

func (s serv) Update(ctx context.Context, user *model.UpdateUserInfo) error {
	err := s.authRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
