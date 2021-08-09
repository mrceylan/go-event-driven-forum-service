package auth

import (
	"context"
	"errors"
	dataaccess "user-service/data-access"
	"user-service/models"
	"user-service/utils"
)

type IAuthService interface {
	GenerateToken(user models.User) (string, error)
	ValidateToken(token string) (string, error)
	Login(ctx context.Context, email string, password string) (string, error)
}

type AuthService struct {
	Database dataaccess.IDatabase
	JwtUtil  *utils.JwtUtil
}

func Activate(db dataaccess.IDatabase, jwtUtil *utils.JwtUtil) IAuthService {
	return &AuthService{db, jwtUtil}
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

func (as *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	result, err := as.Database.GetUserByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	hashedPwd := as.Database.GetUserPasswordById(ctx, result.Id)

	if !utils.ComparePasswords(hashedPwd, password) {
		return "", errors.New("password does not match")
	}

	return as.GenerateToken(result)
}
