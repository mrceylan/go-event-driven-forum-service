package server

import (
	"fmt"
	"forum-service/data-access/interfaces"
	"forum-service/server/controllers"
	"forum-service/server/routes"
	"forum-service/services"

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
	topicService := services.TopicService{
		Repo: repo,
	}

	controller := controllers.Controller{
		TopicService: topicService,
	}
	route := routes.Route{
		Controller: controller,
		App:        app,
	}

	route.TopicRouteMapping()

}
