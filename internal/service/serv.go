package service

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo     repository.Repository
	Validate *validator.Validate
}

func New(repo repository.Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		Validate: validate,
	}
}
