package server

import (
	"github.com/alganbr/kedai-authsvc/configs"
	"github.com/alganbr/kedai-authsvc/internal/clients"
	"github.com/alganbr/kedai-authsvc/internal/controllers"
	"github.com/alganbr/kedai-authsvc/internal/databases"
	"github.com/alganbr/kedai-authsvc/internal/managers"
	"github.com/alganbr/kedai-authsvc/internal/repos"
	"github.com/alganbr/kedai-authsvc/internal/routes"
	"go.uber.org/fx"
)

var controller = fx.Options(
	fx.Provide(controllers.NewHomeController),
	fx.Provide(controllers.NewAuthController),
)

var manager = fx.Options(
	fx.Provide(managers.NewAuthManager),
)

var repo = fx.Options(
	fx.Provide(repos.NewAuthRepository),
)

var database = fx.Options(
	fx.Provide(databases.NewDB),
)

var router = fx.Options(
	fx.Provide(routes.NewRouter),
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewHomeRoutes),
	fx.Provide(routes.NewAuthRoutes),
)

var server = fx.Options(
	fx.Provide(configs.NewConfig),
)

var client = fx.Options(
	fx.Provide(clients.NewUserSvcClient),
)

var Module = fx.Options(
	server,
	client,
	database,
	router,
	controller,
	manager,
	repo,
	fx.Invoke(StartApplication),
)
