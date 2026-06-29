package auth

import (
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
)

type serv struct {
	authRepo  repository.AuthRepository
	redisRepo repository.AuthRepository
	txManager db.TxManager
}

func NewService(authRepo repository.AuthRepository, txManager db.TxManager) service.AuthService {
	return serv{authRepo: authRepo, txManager: txManager}
}
