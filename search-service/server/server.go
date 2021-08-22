package server

import (
	"fmt"
	"log"
	"search-service/server/controllers"
	"search-service/server/routes"
	"search-service/services/search"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port          int
	SearchService search.ISearchService
}

func NewServer(port int, searchsrv search.ISearchService) *Server {
	return &Server{port, searchsrv}
}

func (s *Server) StartServer() error {

	app := fiber.New()
	RegisterRoutes(app, s)
	log.Println("Web server started..")
	err := app.Listen(fmt.Sprintf(":%v", s.Port))

	return err
}

func RegisterRoutes(app *fiber.App, srv *Server) {
	searchController := controllers.NewSearchController(srv.SearchService)

	route := routes.NewRoute(app, searchController)

	route.RouteMapping()
}
