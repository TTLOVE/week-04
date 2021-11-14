//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"week-04/internal/biz"
	// "week-04/internal/conf"
	"week-04/internal/data"
	"week-04/internal/server"
	"week-04/internal/service"

	"github.com/google/wire"
	"github.com/kataras/iris/v12"
)

// initApp init application.
func initApp() (*iris.Application, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

// func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
// 	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet))
// }
