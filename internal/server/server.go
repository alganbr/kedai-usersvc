package server

import (
	"context"
	"github.com/alganbr/kedai-usersvc/configs"
	"github.com/alganbr/kedai-usersvc/internal/databases"
	"github.com/alganbr/kedai-usersvc/internal/routes"
	"go.uber.org/fx"
)

func StartApplication(
	lifecycle fx.Lifecycle,
	cfg *configs.Config,
	router routes.Router,
	routes routes.Routes,
	db *databases.DB,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			db.RunMigration(cfg)
			routes.Setup()
			router.Run(cfg)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer db.Pool.Close()
			return nil
		},
	})
}
