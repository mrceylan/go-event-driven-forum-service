package routes

import (
	"forum-service/server/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App               *fiber.App
	TopicController   controllers.TopicController
	MessageController controllers.MessageController
}

func NewRoute(app *fiber.App,
	topicController controllers.TopicController,
	messageController controllers.MessageController) *Route {
	return &Route{app, topicController, messageController}
}

func (r *Route) RouteMapping() {
	r.topicRouteMapping()
	r.messageRouteMapping()
}

func (r *Route) topicRouteMapping() {

	userApp := r.App.Group("/topic")

	userApp.Post("", r.TopicController.CreateTopic)
	userApp.Get("/:id", r.TopicController.GetTopic)
	userApp.Get("/messages/:id", r.TopicController.GetTopicMessages)
}

func (r *Route) messageRouteMapping() {

	userApp := r.App.Group("/message")

	userApp.Post("", r.MessageController.SaveMessage)
	userApp.Delete("/:id", r.MessageController.DeleteMessage)
	userApp.Get("/:id", r.MessageController.GetMessageById)
}
