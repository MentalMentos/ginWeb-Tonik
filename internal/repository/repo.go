package repository

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/logger"
	"gorm.io/gorm"
)

type RepoImpl struct {
	DB     *gorm.DB
	logger logger.Logger
}

func NewRepo(db *gorm.DB, logger logger.Logger) *RepoImpl {
	return &RepoImpl{
		db,
		logger,
	}
}
