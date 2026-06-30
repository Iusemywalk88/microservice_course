package auth

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
)

func (i *AuthImplementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.Delete(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("Delete called with req: %v", req.GetId())

	return &emptypb.Empty{}, nil
}
