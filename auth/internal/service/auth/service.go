package auth

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
)

type serv struct {
	authRepo repository.AuthRepository
}

func NewService(authRepo repository.AuthRepository) service.AuthService {
	return serv{authRepo: authRepo}
}
