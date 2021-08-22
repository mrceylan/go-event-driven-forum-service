package routes

import (
	"search-service/server/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App              *fiber.App
	SearchController controllers.SearchController
}

func NewRoute(app *fiber.App,
	searchController controllers.SearchController) *Route {
	return &Route{app, searchController}
}

func (r *Route) RouteMapping() {
	r.searchRouteMapping()
}

func (r *Route) searchRouteMapping() {

	userApp := r.App.Group("")

	userApp.Post("saveMessage", r.SearchController.SaveMessage)
	userApp.Post("searchMessages", r.SearchController.SearchMessages)
}
