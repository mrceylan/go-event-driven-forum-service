package controllers

import (
	"user-service/models"

	"github.com/gofiber/fiber/v2"
)

type CreateUser struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc *Controller) CreateUser(ctx *fiber.Ctx) error {
	user := &CreateUser{}

	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := uc.UserService.CreateUser(models.User{
		Email:    user.Email,
		UserName: user.Email,
	}, user.Password)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (uc *Controller) GetUserById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := uc.UserService.GetUserById(id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
