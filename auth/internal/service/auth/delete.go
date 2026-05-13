package auth

import (
	"context"
)

func (s serv) Delete(ctx context.Context, id int64) error {
	err := s.authRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
