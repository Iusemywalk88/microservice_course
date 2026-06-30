package auth

import (
	"context"

	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
)

func (s serv) Update(ctx context.Context, user *model.UpdateUserInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.authRepo.Update(ctx, user)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
