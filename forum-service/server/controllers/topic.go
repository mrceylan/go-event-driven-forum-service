package controllers

import (
	"forum-service/services"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	TopicService services.TopicService
}

type CreateTopic struct {
	Header string `json:"header"`
}

func (uc *Controller) CreateTopic(ctx *fiber.Ctx) error {
	topic := &CreateTopic{}

	err := ctx.BodyParser(&topic)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := uc.TopicService.CreateTopic(topic.Header, "")

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (uc *Controller) GetTopic(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := uc.TopicService.GetTopicById(id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
