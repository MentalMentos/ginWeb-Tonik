package service

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
)

type Service struct {
	*AuthService
}

func New(repo repository.Repository) *Service {
	return &Service{
		NewAuthService(repo),
	}
}
