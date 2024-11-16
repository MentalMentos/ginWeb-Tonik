// service/auth_service.go
package service

import (
	"context"
	"errors"
	_ "github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"

	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	_ "github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req request.RegisterUserRequest) (*model.AuthResponse, error)
	Login(ctx context.Context, req request.LoginRequest) (*model.AuthResponse, error)
}

func (s *Service) Register(ctx context.Context, req request.RegisterUserRequest) (*model.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}
	if err := s.repo.Create(ctx, req) {
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) Login(ctx context.Context, req request.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
