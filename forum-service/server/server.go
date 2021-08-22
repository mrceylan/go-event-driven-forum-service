package server

import (
	"fmt"
	"forum-service/server/controllers"
	"forum-service/server/routes"
	"forum-service/services/message"
	"forum-service/services/topic"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port           int
	TopicService   topic.ITopicService
	MessageService message.IMessageService
}

func NewServer(port int, topicsrv topic.ITopicService, messagesrv message.IMessageService) *Server {
	return &Server{port, topicsrv, messagesrv}
}

func (s *Server) StartServer() error {

	app := fiber.New()
	RegisterRoutes(app, s)
	log.Println("Web server started..")
	err := app.Listen(fmt.Sprintf(":%v", s.Port))

	return err
}

func RegisterRoutes(app *fiber.App, srv *Server) {
	topicController := controllers.NewTopicController(srv.TopicService)
	messageController := controllers.NewMessageController(srv.MessageService)

	route := routes.NewRoute(app, topicController, messageController)

	route.RouteMapping()
}
