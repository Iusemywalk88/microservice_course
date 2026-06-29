package app

import (
	"context"
	"github.com/Iusemywalk88/closer"
	"github.com/Iusemywalk88/microservice_course/auth/internal/api/auth"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/cache"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/cache/redis"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/db/pg"
	"github.com/Iusemywalk88/microservice_course/auth/internal/client/db/transaction"
	"github.com/Iusemywalk88/microservice_course/auth/internal/config"
	"github.com/Iusemywalk88/microservice_course/auth/internal/config/env"
	"github.com/Iusemywalk88/microservice_course/auth/internal/repository"
	authRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/pg"
	redisRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth/redis"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	authService "github.com/Iusemywalk88/microservice_course/auth/internal/service/auth"
	redigo "github.com/gomodule/redigo/redis"
	"log"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool       *redigo.Pool
	redisClient     cache.RedisClient
	redisRepository repository.AuthRepository

	authRepository repository.AuthRepository
	authService    service.AuthService

	authImpl *auth.AuthImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) RedisClient() cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) RedisRepository(ctx context.Context) repository.AuthRepository {
	if s.redisRepository == nil {
		s.redisRepository = redisRepo.NewRepository(
			s.RedisClient(),
			s.RedisConfig().Expiration(),
		)
	}
	return s.redisRepository
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		pgRepo := authRepo.NewRepository(s.DBClient(ctx))
		rediRepo := s.RedisRepository(ctx)

		s.authRepository = repository.NewCachedAuthRepository(pgRepo, rediRepo)
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx), s.TxManager(ctx))
	}

	return s.authService
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.AuthImplementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewAuthImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
