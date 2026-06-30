package app

import (
	"context"
	"github.com/Iusemywalk88/closer"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/api/chat"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/db"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/db/pg"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/client/db/transaction"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/config"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/repository"
	chatRepo "github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
	chatService "github.com/Iusemywalk88/microservice_course/chat-server/internal/service/chat"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	chatRepository repository.ChatRepository
	txManager      db.TxManager

	chatService service.ChatService

	chatImpl *chat.ChatImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatal(err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
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
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.ChatImplementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewChatImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
