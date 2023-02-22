package server

import (
	"github.com/go-kratos/kratos/v2/registry"

	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/util/bootstrap"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	return bootstrap.NewConsulRegistry(conf.Consul.Address, conf.Consul.Scheme, conf.Consul.HealthCheck)
}
