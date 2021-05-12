package routes

func (r Route) AuthRouteMapping() {

	authApp := r.App.Group("/auth")

	authApp.Post("/login", r.Controller.Login)
	authApp.Post("/validateToken", r.Controller.ValidateToken)
}
