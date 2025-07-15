package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	authpb "github.com/bekhuli/notes-service/api/auth"
	"github.com/bekhuli/notes-service/internal/auth"
	"github.com/bekhuli/notes-service/pkg/db"
)

func main() {
	db.Connect()
	defer db.DB.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authRepository := auth.NewUserRepository(db.DB)
	authService := auth.NewAuthService(authRepository)
	authpb.RegisterAuthServiceServer(grpcServer, authService)

	log.Println("AuthService listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
