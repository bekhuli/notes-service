package main

import (
	"context"
	"log"
	"time"

	authpb "github.com/bekhuli/notes-service/api/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	resp, err := client.RegisterUser(ctx, &authpb.RegisterRequest{
		Username: "ds",
		Password: "1234",
	})
	if err != nil {
		log.Fatalf("Register error: %v", err)
	}

	log.Printf("User registered: %+v", resp)
}
