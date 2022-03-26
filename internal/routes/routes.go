package routes

type Routes struct {
	routes []Route
}

type Route interface {
	Setup()
}

func NewRoutes(
	swaggerRoutes SwaggerRoutes,
	homeRoutes HomeRoutes,
	userRoutes UserRoutes,
	passwordRoutes PasswordRoutes,
) Routes {
	return Routes{
		routes: []Route{
			&swaggerRoutes,
			&homeRoutes,
			&userRoutes,
			&passwordRoutes,
		},
	}
}

func (r *Routes) Setup() {
	for _, route := range r.routes {
		route.Setup()
	}
}
