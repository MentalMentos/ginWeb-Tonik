package repository

import (
	"context"
	"errors"

	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, us model.User) (int64, error)
	Update(ctx context.Context, us model.User) (int64, error)
	Delete(ctx context.Context, usId int64) error
	UpdatePassword(ctx context.Context, us model.User, hashPassword string) (model.User, error)
	UpdateIP(ctx context.Context, us model.User, ip string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByID(ctx context.Context, userID int64) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
}

func (r *RepoImpl) Create(ctx context.Context, us model.User) (int64, error) {
	if err := r.DB.WithContext(ctx).Create(&us).Error; err != nil {
		return 0, errors.New("cannot create new user")
	}
	return us.ID, nil
}

func (r *RepoImpl) Update(ctx context.Context, us model.User) (int64, error) {
	updateData := request.UpdateUserRequest{
		Name:  us.Name,
		Email: us.Email,
	}

	if err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", us.ID).Updates(updateData).Error; err != nil {
		return 0, errors.New("cannot update user")
	}
	return us.ID, nil
}

func (r *RepoImpl) Delete(ctx context.Context, usId int64) error {
	if err := r.DB.WithContext(ctx).Delete(&model.User{}, usId).Error; err != nil {
		return errors.New("cannot delete user")
	}
	return nil
}

func (r *RepoImpl) UpdatePassword(ctx context.Context, us model.User, hashPassword string) (model.User, error) {
	updateUser := request.UpdateUserRequest{
		Name:     us.Name,
		Email:    us.Email,
		Password: hashPassword,
	}
	if err := r.DB.WithContext(ctx).Updates(&updateUser).Error; err != nil {
		return model.User{}, errors.New("cannot update password")
	}
	return us, nil
}

func (r *RepoImpl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (r *RepoImpl) GetByID(ctx context.Context, userID int64) (model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (r *RepoImpl) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, errors.New("users not found")
	}
	return users, nil
}

func (r *RepoImpl) UpdateIP(ctx context.Context, us model.User, ip string) (model.User, error) {
	updateUser := request.UpdateUserRequest{
		Name:     us.Name,
		Email:    us.Email,
		Password: us.Password,
		IP:       ip,
	}

	if err := r.DB.WithContext(ctx).Updates(&updateUser).Error; err != nil {
		return us, errors.New("cannot update password")
	}
	return us, nil
}
