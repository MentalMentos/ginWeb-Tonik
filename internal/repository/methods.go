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
	err := r.DB.Create(us).WithContext(ctx)
	if err == nil {
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
	err := r.DB.Model(&model.User{}).Updates(updateUser).WithContext(ctx)
	if err == nil {
		return us.ID, errors.New("id  Cannot update user")
	}
	return updateUser.Id, nil
}

func (r *RepoImpl) Delete(ctx context.Context, usId int64) (int64, error) {
	err := r.DB.Delete(usId).WithContext(ctx)
	if err == nil {
		return usId, errors.New("id  Cannot delete user")
	}
	return usId, nil
}

func (r *RepoImpl) GetByID(ctx context.Context, id int64) (model.User, error) {
	var us model.User
	err := r.DB.Find(&us, id).WithContext(ctx)
	if err == nil {
		return us, errors.New("user is not found")
	}
	return us, nil
}

func (r *RepoImpl) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users)
	if err == nil {
		return users, nil
	}
	return users, nil
}
