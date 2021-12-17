package usercase

import (
	"BazarMessenger/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func (this *userUsecase) Create(user *domain.User, context context.Context) error {
	panic("implement me")
}

func (this *userUsecase) Update(user *domain.User, context context.Context) error {
	panic("implement me")
}

func (this *userUsecase) GetById(id string, context context.Context) (*domain.User, error) {
	filter := bson.D{primitive.E{Key: "id", Value: id}}

	user, err := this.userRepo.GetById(filter, context)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (this *userUsecase) GetByUsername(userName string, context context.Context) (*domain.User, error) {
	filter := bson.D{primitive.E{Key: "user_name", Value: userName}}
	user, err := this.userRepo.GetByUsername(filter, context)
	fmt.Println("Error: ", err)

	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func NewUserUsecase(userrep domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userrep,
	}
}
