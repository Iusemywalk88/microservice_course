package converter

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/pg/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		UserRole:  model.Role(user.UserRole),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToRepoFromUser(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		UserRole:  modelRepo.Role(user.UserRole),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToRepoFromServiceUpdate(user *model.UpdateUserInfo) *modelRepo.UpdateUserInfo {
	return &modelRepo.UpdateUserInfo{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UserRole: modelRepo.Role(user.UserRole),
	}
}
