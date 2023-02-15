package server

import (
	"github.com/go-kratos/kratos/v2/registry"

	"kratos-blog/app/content/service/internal/conf"
	"kratos-blog/pkg/util/bootstrap"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	return bootstrap.NewConsulRegistrar(conf.Consul.Address, conf.Consul.Scheme, conf.Consul.HealthCheck)
}
