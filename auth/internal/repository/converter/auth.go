package converter

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/model"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAuthFromRepo(auth *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if auth.UpdatedAt.Valid {
		updatedAt = timestamppb.New(auth.UpdatedAt.Time)
	}
	return &desc.GetResponse{
		Id:        auth.ID,
		Name:      auth.Name,
		Email:     auth.Email,
		Role:      desc.Role(auth.UserRole),
		CreatedAt: timestamppb.New(auth.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
