package bootstrap

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kratosRest "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/gorilla/handlers"

	"kratos-blog/gen/api/go/common/conf"
)

// CreateRestServer 创建REST服务端
func CreateRestServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosRest.Server {
	var opts = []kratosRest.ServerOption{
		kratosRest.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Rest.Headers),
			handlers.AllowedMethods(cfg.Server.Rest.Methods),
			handlers.AllowedOrigins(cfg.Server.Rest.Origins),
		)),
	}

	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, kratosRest.Middleware(ms...))

	if cfg.Server.Rest.Network != "" {
		opts = append(opts, kratosRest.Network(cfg.Server.Rest.Network))
	}
	if cfg.Server.Rest.Addr != "" {
		opts = append(opts, kratosRest.Address(cfg.Server.Rest.Addr))
	}
	if cfg.Server.Rest.Timeout != nil {
		opts = append(opts, kratosRest.Timeout(cfg.Server.Rest.Timeout.AsDuration()))
	}

	return kratosRest.NewServer(opts...)
}
