package server

import (
	"fmt"
	"log"
	"user-service/server/controllers"
	"user-service/server/routes"
	"user-service/services/auth"
	"user-service/services/user"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port        int
	UserService user.IUserService
	AuthService auth.IAuthService
}

func NewServer(port int, usrsrv user.IUserService, authsrv auth.IAuthService) *Server {
	return &Server{port, usrsrv, authsrv}
}

func (s *Server) StartServer() error {

	app := fiber.New()
	RegisterRoutes(app, s)
	log.Println("Web server started..")
	err := app.Listen(fmt.Sprintf(":%v", s.Port))

	return err
}

func RegisterRoutes(app *fiber.App, srv *Server) {
	authController := controllers.NewAuthController(srv.AuthService)

	userController := controllers.NewUserController(srv.UserService)

	routes := routes.NewRoute(app, userController, authController)

	routes.RouteMapping()
}
