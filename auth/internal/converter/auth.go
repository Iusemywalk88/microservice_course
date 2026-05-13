package converter

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/model"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.UserRole),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToServiceFromUserCreate(desc *desc.CreateRequest) *model.User {
	return &model.User{
		Name:     desc.Name,
		Email:    desc.Email,
		Password: desc.Password,
		UserRole: model.Role(desc.Role),
	}
}

func ToServiceFromUserUpdate(desc *desc.UpdateRequest) *model.UpdateUserInfo {
	var name *string
	if desc.GetName() != nil {
		val := desc.GetName().GetValue()
		name = &val
	}

	var email *string
	if desc.GetEmail() != nil {
		val := desc.GetEmail().GetValue()
		email = &val
	}

	return &model.UpdateUserInfo{
		ID:       desc.Id,
		Name:     name,
		Email:    email,
		UserRole: model.Role(desc.Role),
	}
}
