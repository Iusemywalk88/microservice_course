package auth

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
