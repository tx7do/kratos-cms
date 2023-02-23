package server

import (
	"github.com/go-kratos/kratos/v2/registry"

	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/bootstrap"
)

func NewRegistrar(cfg *conf.Registry) registry.Registrar {
	return bootstrap.NewConsulRegistry(cfg)
}
