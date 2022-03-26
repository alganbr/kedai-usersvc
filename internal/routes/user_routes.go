package routes

import "github.com/alganbr/kedai-usersvc/internal/controllers"

type UserRoutes struct {
	router         Router
	userController controllers.IUserController
}

func NewUserRoutes(router Router, userController controllers.IUserController) UserRoutes {
	return UserRoutes{
		router:         router,
		userController: userController,
	}
}

func (r *UserRoutes) Setup() {
	userGroup := r.router.Path.Group("/user")
	userGroup.GET("/:id", r.userController.Get)
	userGroup.POST("", r.userController.Create)
	userGroup.PUT("/:id", r.userController.Update)
	userGroup.PATCH("/:id", r.userController.Patch)
}
