package repository

import (
	"context"
	"errors"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
)

type Repository interface {
	Create(ctx context.Context, us model.User) (int64, error)
	Update(ctx context.Context, us model.User) (int64, error)
	Delete(ctx context.Context, usId int64) (int64, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
}

func (r *RepoImpl) Create(ctx context.Context, us model.User) (int64, error) {
	tx := r.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := r.DB.Create(&us).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return us.ID, errors.New("id  Cannot create new user")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return us.ID, errors.New("id  Cannot create new user")
	}
	return us.ID, nil

}

func (r *RepoImpl) Update(ctx context.Context, us model.User) (int64, error) {
	var updateUser = request.UpdateUserRequest{
		us.ID,
		us.Name,
		us.Email,
	}
	tx := r.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := r.DB.Model(&model.User{}).Updates(&updateUser).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return us.ID, errors.New("id  Cannot update user")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return us.ID, errors.New("id  Cannot update user")
	}
	return updateUser.Id, nil
}

func (r *RepoImpl) Delete(ctx context.Context, usId int64) (int64, error) {
	tx := r.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := r.DB.Delete(&usId).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return usId, errors.New("id  Cannot delete user")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return usId, errors.New("id  Cannot delete user")
	}
	return usId, nil
}

func (r *RepoImpl) GetByID(ctx context.Context, id int64) (model.User, error) {
	var us model.User

	tx := r.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := r.DB.Find(&us, id).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return us, errors.New("user is not found")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return us, errors.New("user is not found")
	}
	return us, nil
}

func (r *RepoImpl) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	tx := r.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := r.DB.Find(&users).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return users, errors.New("users are not found")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return users, errors.New("users are not found")
	}
	return users, nil
}
