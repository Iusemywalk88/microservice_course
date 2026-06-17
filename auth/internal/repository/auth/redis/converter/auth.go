package converter

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/pg/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/redis/model"
)

func updatedAtToInt(auth *model.User) *int64 {
	if auth.UpdatedAt.Valid == true {
		unix := auth.UpdatedAt.Time.Unix()
		return &unix
	}
	return nil
}

func ToRepoFromUser(auth *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        auth.ID,
		Name:      auth.Name,
		Email:     auth.Email,
		Password:  auth.Password,
		UserRole:  modelRepo.Role(auth.UserRole),
		CreatedAt: auth.CreatedAt.Unix(),
		UpdatedAt: updatedAtToInt(auth),
	}
}
