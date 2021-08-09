package controllers

import (
	"user-service/models"
	"user-service/services/user"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService user.IUserService
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
	type CreateUser struct {
		UserName string `json:"userName"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	user := &CreateUser{}

	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := uc.UserService.CreateUser(ctx.Context(), models.User{
		Email:    user.Email,
		UserName: user.Email,
	}, user.Password)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (uc *UserController) GetUserById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := uc.UserService.GetUserById(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
