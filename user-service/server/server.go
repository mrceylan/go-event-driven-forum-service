package server

import (
	"fmt"
	"user-service/data-access/interfaces"
	"user-service/server/controllers"
	"user-service/server/routes"
	"user-service/services"
	"user-service/utils"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port       int
	Repository interfaces.Repository
}

func (s *Server) CreateServer() {
	app := fiber.New()

	RegisterRoutes(app, s.Repository)

	app.Listen(fmt.Sprintf(":%v", s.Port))
}

func RegisterRoutes(app *fiber.App, repo interfaces.Repository) {
	userService := services.UserService{
		Repo: repo,
	}
	authService := services.AuthService{
		JwtUtil: utils.JwtUtil{
			SecretKey:       "emre",
			ExpirationHours: 12,
		},
	}

	controller := controllers.Controller{
		UserService: userService,
		AuthService: authService,
	}
	route := routes.Route{
		Controller: controller,
		App:        app,
	}

	route.UserRouteMapping()
	route.AuthRouteMapping()
}
