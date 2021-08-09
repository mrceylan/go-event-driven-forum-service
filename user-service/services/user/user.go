package user

import (
	"context"
	dataaccess "user-service/data-access"
	"user-service/models"
	"user-service/utils"
)

type IUserService interface {
	CreateUser(ctx context.Context, model models.User, password string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type UserService struct {
	Database dataaccess.IDatabase
}

func Activate(db dataaccess.IDatabase) IUserService {
	return &UserService{Database: db}
}

func (us *UserService) CreateUser(ctx context.Context, model models.User, password string) (models.User, error) {

	hashedPwd, err := utils.HashAndSaltPassword(password)

	if err != nil {
		return models.User{}, err
	}

	result, err := us.Database.SaveUser(ctx, model, hashedPwd)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (us *UserService) GetUserById(ctx context.Context, id string) (models.User, error) {

	result, err := us.Database.GetUserById(ctx, id)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	result, err := us.Database.GetUserByEmail(ctx, email)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}
