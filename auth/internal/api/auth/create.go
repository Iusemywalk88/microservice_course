package auth

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/converter"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *AuthImplementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.PasswordConfirm != req.Password {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	id, err := i.authService.Create(ctx, converter.ToServiceFromUserCreate(req))
	if err != nil {
		return nil, err
	}

	log.Printf("Create called with id: %v", id)

	return &desc.CreateResponse{Id: id}, nil
}
