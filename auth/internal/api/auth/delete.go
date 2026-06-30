package auth

import (
	"context"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (i *AuthImplementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("Delete called with req: %v", req.GetId())

	return &emptypb.Empty{}, nil
}
