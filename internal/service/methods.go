package service

import (
	"context"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/request"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/data/response"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"log"
)

type UserService interface {
	Create(ctx context.Context, user request.CreateUserRequest)
	Update(ctx context.Context, user request.UpdateUserRequest)
	Delete(ctx context.Context, userId int64)
	FindById(ctx context.Context, userId int64) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}

func (serv *Service) Create(ctx context.Context, user request.CreateUserRequest) {
	err := serv.Validate.Struct(user)
	if err != nil {
		log.Fatalf("Cannot to create user in service: ", err)
	}
	usModel := model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	serv.Repository.Create(ctx, usModel)
}

func (serv *Service) Update(ctx context.Context, user request.UpdateUserRequest) {
	data, err := serv.Repository.GetByID(ctx, user.Id)
	if err != nil {
		log.Fatalf("Cannot to update user in service: ", err)
	}
	data.Name = user.Name
	data.Email = user.Email
	serv.Repository.Update(ctx, data)
}

func (serv *Service) Delete(ctx context.Context, userId int64) {
	serv.Repository.Delete(ctx, userId)
}

func (serv *Service) FindById(ctx context.Context, userId int64) response.UserResponse {
	data, err := serv.Repository.GetByID(ctx, userId)
	if err != nil {
		log.Fatalf("Cannot to find user in service: ", err)
	}

	userResponse := response.UserResponse{
		Id:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
	return userResponse
}

func (serv *Service) FindAll(ctx context.Context) []response.UserResponse {
	result, err := serv.Repository.GetAll(ctx)
	if err != nil {
		log.Fatalf("Cannot to find user in service: ", err)
	}

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:   value.ID,
			Name: value.Name,
		}
		users = append(users, user)
	}

	return users
}
