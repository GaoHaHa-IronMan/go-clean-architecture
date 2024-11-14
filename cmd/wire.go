//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"go-clean-architecture/api"
	"go-clean-architecture/api/handlers"
	"go-clean-architecture/app"
	"go-clean-architecture/repo"
	"go-clean-architecture/service"

	"github.com/google/wire"
)

func InitServer() *app.Server {
	wire.Build(
		app.InitDB,
		repo.NewMysqlArticleRepository,
		service.NewArticleService,
		handlers.NewArticleHandler,
		api.NewRouter,
		app.NewServer,
		app.NewGinEngine)
	return &app.Server{}
}
