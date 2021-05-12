package services

import (
	"user-service/models"
	"user-service/utils"
)

type AuthService struct {
	JwtUtil utils.JwtUtil
}

func (as *AuthService) GenerateToken(user models.User) (string, error) {

	token, err := as.JwtUtil.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) ValidateToken(token string) (string, error) {

	id, err := as.JwtUtil.ValidateToken(token)

	if err != nil {
		return "", err
	}

	return id, nil
}
