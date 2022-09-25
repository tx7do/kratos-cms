//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-blog/app/content/service/internal/biz"
	"kratos-blog/app/content/service/internal/conf"
	"kratos-blog/app/content/service/internal/data"
	"kratos-blog/app/content/service/internal/server"
	"kratos-blog/app/content/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
