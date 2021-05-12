package controllers

import (
	"user-service/constants"
	"user-service/services"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	UserService services.UserService
	AuthService services.AuthService
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	AccessToken string `json:"access_token"`
}

type ValidateTokenResult struct {
	Id string `json:"id"`
}

func (ac *Controller) Login(ctx *fiber.Ctx) error {

	login := &Login{}
	err := ctx.BodyParser(&login)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := ac.UserService.CheckUserPassword(login.Email, login.Password)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err)
	}

	token, err := ac.AuthService.GenerateToken(user)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(LoginResult{
		AccessToken: token,
	})

}

func (ac *Controller) ValidateToken(ctx *fiber.Ctx) error {

	token := ctx.Get(constants.TOKEN_HEADER_NAME)

	id, err := ac.AuthService.ValidateToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(ValidateTokenResult{
		Id: id,
	})

}
