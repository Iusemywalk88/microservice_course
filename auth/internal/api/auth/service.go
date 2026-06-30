package auth

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
)

type AuthImplementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func NewAuthImplementation(authService service.AuthService) *AuthImplementation {
	return &AuthImplementation{authService: authService}
}
