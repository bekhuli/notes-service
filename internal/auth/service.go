package auth

import (
	"context"

	authpb "github.com/bekhuli/notes-service/api/auth"
	pb "github.com/bekhuli/notes-service/api/auth"
	"github.com/bekhuli/notes-service/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	repo Repository
}

func NewAuthService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	id, err := s.repo.RegisterUser(ctx, req.Username, string(hash))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id, Username: req.Username}, nil
}

func (s *Service) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	id, err := s.repo.RegisterUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token := auth.GenerateJWT(id)

	return &pb.AuthResponse{
		Token: token,
		User:  &pb.UserResponse{Id: id, Username: req.Username},
	}, nil
}
