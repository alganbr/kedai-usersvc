package server

import (
	"github.com/alganbr/kedai-usersvc/configs"
	"github.com/alganbr/kedai-usersvc/internal/controllers"
	"github.com/alganbr/kedai-usersvc/internal/databases"
	"github.com/alganbr/kedai-usersvc/internal/managers"
	"github.com/alganbr/kedai-usersvc/internal/repos"
	"github.com/alganbr/kedai-usersvc/internal/routes"
	"github.com/alganbr/kedai-usersvc/internal/utils/logger"
	"go.uber.org/fx"
)

var controller = fx.Options(
	fx.Provide(controllers.NewHomeController),
	fx.Provide(controllers.NewUserController),
	fx.Provide(controllers.NewPasswordController),
)

var manager = fx.Options(
	fx.Provide(managers.NewUserManager),
	fx.Provide(managers.NewPasswordManager),
)

var repo = fx.Options(
	fx.Provide(repos.NewUserRepository),
	fx.Provide(repos.NewPasswordRepository),
)

var database = fx.Options(
	fx.Provide(databases.NewDB),
)

var router = fx.Options(
	fx.Provide(routes.NewRouter),
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewHomeRoutes),
	fx.Provide(routes.NewUserRoutes),
	fx.Provide(routes.NewPasswordRoutes),
)

var server = fx.Options(
	fx.Provide(configs.NewConfig),
	fx.Provide(logger.NewLogger),
)

var Module = fx.Options(
	server,
	database,
	router,
	controller,
	manager,
	repo,
	fx.Invoke(StartApplication),
)
