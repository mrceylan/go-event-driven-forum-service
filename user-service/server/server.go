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
	defer log.Println("Web server started")

	app := fiber.New()
	RegisterRoutes(app, s)
	err := app.Listen(fmt.Sprintf(":%v", s.Port))

	return err
}

func RegisterRoutes(app *fiber.App, srv *Server) {
	authController := controllers.AuthController{
		AuthService: srv.AuthService,
	}

	userController := controllers.UserController{
		UserService: srv.UserService,
	}

	routes := routes.Route{
		UserController: userController,
		AuthController: authController,
		App:            app,
	}

	routes.RouteMapping()
}
