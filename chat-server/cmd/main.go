package main

import (
	"context"
	"flag"
	config "github.com/Iusemywalk88/microservice_course/chat-server/internal/config"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/converter"
	repoChat "github.com/Iusemywalk88/microservice_course/chat-server/internal/repository/chat"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service"
	"github.com/Iusemywalk88/microservice_course/chat-server/internal/service/chat"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../.env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatModel, err := converter.ToServiceFromDesc(req)
	if err != nil {
		return nil, err
	}

	id, err := s.chatService.Create(ctx, chatModel)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete request: %v", req.GetId())

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("SendMessage request: %s, %s, %v", req.GetFrom(), req.GetText(), req.GetTimestamp())

	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	chatRepo := repoChat.NewRepository(pool)
	chatService := chat.NewService(chatRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{chatService: chatService})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
