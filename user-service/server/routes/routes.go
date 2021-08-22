package routes

import (
	"user-service/server/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App            *fiber.App
	UserController controllers.UserController
	AuthController controllers.AuthController
}

func NewRoute(app *fiber.App, userController controllers.UserController, authController controllers.AuthController) *Route {
	return &Route{app, userController, authController}
}

func (r *Route) RouteMapping() {
	r.userRouteMapping()
	r.authRouteMapping()
}

func (r *Route) userRouteMapping() {

	userApp := r.App.Group("/user")

	userApp.Post("", r.UserController.CreateUser)
	userApp.Get("/:id", r.UserController.GetUserById)
}

func (r *Route) authRouteMapping() {

	authApp := r.App.Group("/auth")

	authApp.Post("/login", r.AuthController.Login)
	authApp.Post("/validateToken", r.AuthController.ValidateToken)
}
