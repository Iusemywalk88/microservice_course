package converter

import (
	"database/sql"
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	modelRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/redis/model"
	"time"
)

func updatedAtToInt(user *model.User) *int64 {
	if user.UpdatedAt.Valid == true {
		unix := user.UpdatedAt.Time.Unix()
		return &unix
	}
	return nil
}

func updatedAtToSqlNullTime(user *modelRepo.User) sql.NullTime {
	var updatedAt sql.NullTime
	if user.UpdatedAt != nil {
		updatedAt = sql.NullTime{
			Time:  time.Unix(*user.UpdatedAt, 0),
			Valid: true,
		}
		return updatedAt
	}
	return updatedAt
}

func ToRepoFromUser(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UserRole:  modelRepo.Role(user.UserRole),
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: updatedAtToInt(user),
	}
}

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UserRole:  model.Role(user.UserRole),
		CreatedAt: time.Unix(user.CreatedAt, 0),
		UpdatedAt: updatedAtToSqlNullTime(user),
	}
}
