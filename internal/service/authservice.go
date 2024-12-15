package service

import (
	"context"
	"errors"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	_ "github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/helpers"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/logger"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/utils"

	_ "github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Register(ctx context.Context, req request.RegisterUserRequest, logger logger.Logger) (*model.AuthResponse, error)
	Login(ctx context.Context, req request.LoginRequest, logger logger.Logger) (*model.AuthResponse, error)
	GetAccessToken(ctx context.Context, refreshToken string, logger logger.Logger) (*response.AuthResponse, error)
	UpdatePassword(ctx context.Context, req request.UpdateUserRequest, logger logger.Logger) (*response.UpdatePasswordResponse, error)
}

type AuthService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewAuthService(repo repository.Repository, logger logger.Logger) *AuthService {
	return &AuthService{
		repo,
		logger,
	}
}

func (s *AuthService) Register(ctx context.Context, req request.RegisterUserRequest, logger logger.Logger) (*model.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToHashPass)
		return nil, err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
		IP:       req.IP,
	}

	_, err = s.repo.Create(ctx, user, logger)
	if err != nil {
		logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToCreateUser)
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		logger.Fatal("[ SERVICE_REGISTER ]", helpers.FailedToGenJWT)
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req request.LoginRequest, logger logger.Logger) (*model.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email, logger)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.IP != req.IP {
		_, err := s.repo.UpdateIP(ctx, user, req.IP, logger)
		if err != nil {
			return nil, errors.New("cannot update ip with login")
		}
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

func (s *AuthService) UpdatePassword(ctx context.Context, req request.UpdateUserRequest, logger logger.Logger) (*response.UpdatePasswordResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email, logger)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.IP != req.IP {
		_, err := s.repo.UpdateIP(ctx, user, req.IP, logger)
		if err != nil {
			return nil, errors.New("cannot update ip with login")
		}
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
func (s *AuthService) GetAccessToken(ctx context.Context, refreshToken string, logger logger.Logger) (*response.AuthResponse, error) {
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
