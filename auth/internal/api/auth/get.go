package auth

import (
	"context"
	"github.com/Iusemywalk88/microservice_course/auth/internal/converter"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %d", userObj.ID, userObj.Name, userObj.Email, userObj.UserRole)

	return converter.ToUserFromService(*userObj), nil
}
