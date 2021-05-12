package services

import (
	"errors"
	"user-service/data-access/interfaces"
	"user-service/models"
	"user-service/utils"
)

type UserService struct {
	Repo interfaces.Repository
}

func (us UserService) CreateUser(model models.User, password string) (models.User, error) {

	hashedPwd, err := utils.HashAndSaltPassword(password)

	if err != nil {
		return models.User{}, err
	}

	result, err := us.Repo.SaveUser(model, hashedPwd)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (us UserService) GetUserById(id string) (models.User, error) {

	result, err := us.Repo.GetUserById(id)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (us UserService) GetUserByEmail(email string) (models.User, error) {

	result, err := us.Repo.GetUserByEmail(email)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (us UserService) CheckUserPassword(email string, password string) (models.User, error) {

	result, err := us.Repo.GetUserByEmail(email)

	if err != nil {
		return models.User{}, err
	}

	hashedPwd := us.Repo.GetUserPasswordById(result.Id)

	if !utils.ComparePasswords(hashedPwd, password) {
		return models.User{}, errors.New("password does not match")
	}

	return result, nil
}
