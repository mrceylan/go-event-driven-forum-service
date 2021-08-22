package controllers

import (
	"search-service/models"
	"search-service/services/search"

	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	SearchService search.ISearchService
}

func NewSearchController(searchService search.ISearchService) SearchController {
	return SearchController{searchService}
}

func (ss *SearchController) SaveMessage(ctx *fiber.Ctx) error {
	type SaveMessage struct {
		Id      string `json:"id"`
		Message string `json:"message"`
		TopicId string `json:"topicId"`
	}

	message := &SaveMessage{}

	err := ctx.BodyParser(&message)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	err = ss.SearchService.SaveMessage(ctx.Context(),
		models.Message{
			Id:      message.Id,
			Message: message.Message,
			TopicId: message.TopicId,
		})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON("OK")
}

func (ss *SearchController) SearchMessages(ctx *fiber.Ctx) error {
	type SearchMessages struct {
		SearchString string `json:"searchString"`
	}

	searchMessages := &SearchMessages{}

	err := ctx.BodyParser(&searchMessages)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}

	result, err := ss.SearchService.SearchMessages(ctx.Context(), searchMessages.SearchString)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}
