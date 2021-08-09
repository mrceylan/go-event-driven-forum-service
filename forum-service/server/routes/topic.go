package routes

import (
	"forum-service/server/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Controller controllers.Controller
	App        *fiber.App
}

func (r Route) TopicRouteMapping() {

	userApp := r.App.Group("/topic")

	userApp.Post("", r.Controller.CreateTopic)
	userApp.Get("/:id", r.Controller.GetTopic)
}
