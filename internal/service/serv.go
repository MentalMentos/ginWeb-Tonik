package service

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	Repository repository.Repository
	Validate   *validator.Validate
}

func New(repository repository.Repository, validate *validator.Validate) *Service {
	return &Service{
		Repository: repository,
		Validate:   validate,
	}
}
