package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtil struct {
	SecretKey         string
	ExpirationMinutes int
}

func NewJwtUtil(secretKey string, expirationMinutes int) *JwtUtil {
	return &JwtUtil{secretKey, expirationMinutes}
}

func (j *JwtUtil) GenerateToken(id string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
		Id:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JwtUtil) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return "", err
	}

	return claims.Id, nil

}
