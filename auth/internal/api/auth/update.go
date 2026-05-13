package auth

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/converter"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	updateObj := converter.ToServiceFromUserUpdate(req)

	err := i.authService.Update(ctx, updateObj)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("Update called with req: %v", req.GetId(), req.GetName().GetValue())

	return &emptypb.Empty{}, nil
}
