package routes

import (
	"user-service/server/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Controller controllers.Controller
	App        *fiber.App
}

func (r Route) UserRouteMapping() {

	userApp := r.App.Group("/user")

	userApp.Post("", r.Controller.CreateUser)
	userApp.Get("/:id", r.Controller.GetUserById)
}
