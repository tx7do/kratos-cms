package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-blog/app/content/service/internal/service"
	"kratos-blog/gen/api/go/common/conf"
	v1 "kratos-blog/gen/api/go/content/service/v1"
	"kratos-blog/pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	tagSvc *service.TagService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))

	v1.RegisterPostServiceServer(srv, postSvc)
	v1.RegisterLinkServiceServer(srv, linkSvc)
	v1.RegisterCategoryServiceServer(srv, cateSvc)
	v1.RegisterTagServiceServer(srv, tagSvc)

	return srv
}
