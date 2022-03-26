package routes

import "github.com/alganbr/kedai-usersvc/internal/controllers"

type PasswordRoutes struct {
	router             Router
	passwordController controllers.IPasswordController
}

func NewPasswordRoutes(router Router, passwordController controllers.IPasswordController) PasswordRoutes {
	return PasswordRoutes{
		router:             router,
		passwordController: passwordController,
	}
}

func (r *PasswordRoutes) Setup() {
	passwordGroup := r.router.Path.Group("/password")
	passwordGroup.PUT("", r.passwordController.Update)
	passwordGroup.POST("/validate", r.passwordController.Validate)
}
