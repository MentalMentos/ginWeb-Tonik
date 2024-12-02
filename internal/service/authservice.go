package service

import (
	"context"
	"errors"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
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
	GetAccessToken(ctx context.Context, req request.LoginRequest) (*model.AuthResponse, error)
	UpdatePassword(ctx context.Context, req request.UpdateUserRequest) (*response.UpdatePasswordResponse, error)
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
	_, err = s.repo.Create(ctx, user)
	if err != nil {
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

func (s *Service) Login(ctx context.Context, req request.LoginRequest) (*model.AuthResponse, error) {
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

func (s *Service) UpdatePassword(ctx context.Context, req request.UpdateUserRequest) (*response.UpdatePasswordResponse, error) {
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
	return &response.UpdatePasswordResponse{
		accessToken,
		refreshToken,
		req.Name,
	}, nil
}

// Метод для обновления access token
func (s *Service) GetAccessToken(ctx context.Context, refreshToken string) (*response.AuthResponse, error) {
	// Валидация refresh token
	claims, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Генерация нового набора токенов
	newAccessToken, newRefreshToken, err := utils.GenerateJWT(claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
