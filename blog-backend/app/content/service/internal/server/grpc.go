package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "kratos-blog/gen/api/go/content/service/v1"

	"kratos-blog/app/content/service/internal/service"
	"kratos-blog/gen/api/go/common/conf"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	tagSvc *service.TagService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if cfg.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(cfg.Server.Grpc.Network))
	}
	if cfg.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(cfg.Server.Grpc.Addr))
	}
	if cfg.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(cfg.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterPostServiceServer(srv, postSvc)
	v1.RegisterLinkServiceServer(srv, linkSvc)
	v1.RegisterCategoryServiceServer(srv, cateSvc)
	v1.RegisterTagServiceServer(srv, tagSvc)

	return srv
}
