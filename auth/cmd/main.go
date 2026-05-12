package main

import (
	"context"
	"flag"
	"github.com/Iusemywalk88/microservice_course/auth/internal/config"
	"github.com/Iusemywalk88/microservice_course/auth/internal/converter"
	authRepo "github.com/Iusemywalk88/microservice_course/auth/internal/repository/auth"
	"github.com/Iusemywalk88/microservice_course/auth/internal/service"
	authService "github.com/Iusemywalk88/microservice_course/auth/internal/service/auth"
	desc "github.com/Iusemywalk88/microservice_course/auth/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../.env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := s.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %d", userObj.ID, userObj.Name, userObj.Email, userObj.UserRole)

	return converter.ToUserFromService(*userObj), nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.PasswordConfirm != req.Password {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	id, err := s.authService.Create(ctx, converter.ToServiceFromUserCreate(req))
	if err != nil {
		return nil, err
	}

	log.Printf("Create called with id: %v", id)

	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update called with req: %v", req.GetId(), req.GetName().GetValue())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete called with req: %v", req.GetId())

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

	authRepo := authRepo.NewRepository(pool)
	authSrv := authService.NewService(authRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{authService: authSrv})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
