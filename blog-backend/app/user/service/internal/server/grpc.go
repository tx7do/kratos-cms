package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-blog/app/user/service/internal/service"
	"kratos-blog/gen/api/go/common/conf"
	v1 "kratos-blog/gen/api/go/user/service/v1"
	"kratos-blog/pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))

	v1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
