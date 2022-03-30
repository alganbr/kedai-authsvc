package server

import (
	"context"
	"github.com/alganbr/kedai-authsvc/configs"
	"github.com/alganbr/kedai-authsvc/internal/databases"
	"github.com/alganbr/kedai-authsvc/internal/routes"
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
			return nil
		},
	})
}
