package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-blog/app/comment/service/internal/service"
	v1 "kratos-blog/gen/api/go/comment/service/v1"
	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	commentService *service.CommentService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logger)

	v1.RegisterCommentServiceServer(srv, commentService)

	return srv
}
