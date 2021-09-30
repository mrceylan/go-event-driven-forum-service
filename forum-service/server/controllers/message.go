package controllers

import (
	"forum-service/models"
	"forum-service/services/message"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MessageController struct {
	MessageService message.IMessageService
}

func NewMessageController(messageService message.IMessageService) MessageController {
	return MessageController{messageService}
}

func (mc *MessageController) SaveMessage(ctx *fiber.Ctx) error {
	type SaveMessage struct {
		Message string `json:"message"`
		TopicId string `json:"topicId"`
		UserId  string `json:"userId"`
	}

	message := &SaveMessage{}

	err := ctx.BodyParser(&message)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := mc.MessageService.SaveMessage(ctx.Context(),
		models.Message{
			Message:    message.Message,
			TopicId:    message.TopicId,
			CreateDate: time.Now(),
			CreatedBy:  message.UserId,
		})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (mc *MessageController) DeleteMessage(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	err := mc.MessageService.DeleteMessage(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("OK")
}

func (mc *MessageController) GetMessageById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := mc.MessageService.GetMessageById(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
