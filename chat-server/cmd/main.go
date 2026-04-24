package main

import (
	"context"
	"flag"
	"github.com/Iusemywalk88/microservice_course/chat-server/config"
	desc "github.com/Iusemywalk88/microservice_course/chat-server/pkg/chat_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v6"
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
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create: %v", req.GetUsernames())

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
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

	// Делаем запрос на вставку записи в таблицу note
	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(gofakeit.Name()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatID int
	err = pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert note: %v", err)
	}

	log.Printf("inserted note with id: %d", chatID)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
