//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"

	"github.com/tx7do/kratos-blog/blog-backend/app/file/service/internal/biz"
	"github.com/tx7do/kratos-blog/blog-backend/app/file/service/internal/conf"
	"github.com/tx7do/kratos-blog/blog-backend/app/file/service/internal/data"
	"github.com/tx7do/kratos-blog/blog-backend/app/file/service/internal/server"
	"github.com/tx7do/kratos-blog/blog-backend/app/file/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
