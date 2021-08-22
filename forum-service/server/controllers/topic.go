package controllers

import (
	"forum-service/models"
	"forum-service/services/topic"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TopicController struct {
	TopicService topic.ITopicService
}

func NewTopicController(topicService topic.ITopicService) TopicController {
	return TopicController{topicService}
}

func (tc *TopicController) CreateTopic(ctx *fiber.Ctx) error {
	type CreateTopic struct {
		Header string `json:"header"`
		UserId string `json:"userId"`
	}

	topic := &CreateTopic{}

	err := ctx.BodyParser(&topic)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := tc.TopicService.CreateTopic(ctx.Context(),
		models.Topic{
			Header:     topic.Header,
			CreateDate: time.Now(),
			CreatedBy:  topic.UserId,
		})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (tc *TopicController) GetTopic(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := tc.TopicService.GetTopicById(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (tc *TopicController) GetTopicMessages(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	result, err := tc.TopicService.GetTopicMessages(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
