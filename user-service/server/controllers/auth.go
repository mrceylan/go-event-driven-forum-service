package controllers

import (
	"user-service/constants"
	"user-service/services/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService auth.IAuthService
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {

	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResult struct {
		AccessToken string `json:"access_token"`
	}

	login := &Login{}
	err := ctx.BodyParser(&login)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	token, err := ac.AuthService.Login(ctx.Context(), login.Email, login.Password)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(LoginResult{
		AccessToken: token,
	})

}

func (ac *AuthController) ValidateToken(ctx *fiber.Ctx) error {

	type ValidateTokenResult struct {
		Id string `json:"id"`
	}

	token := ctx.Get(constants.TOKEN_HEADER_NAME)

	id, err := ac.AuthService.ValidateToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(ValidateTokenResult{
		Id: id,
	})

}
