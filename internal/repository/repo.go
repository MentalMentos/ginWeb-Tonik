package repository

import "gorm.io/gorm"

type RepoImpl struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *RepoImpl {
	return &RepoImpl{db}
}
