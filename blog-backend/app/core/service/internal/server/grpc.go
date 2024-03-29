package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	bootstrap "github.com/tx7do/kratos-bootstrap"
	conf "github.com/tx7do/kratos-bootstrap/gen/api/go/conf/v1"

	"kratos-cms/app/core/service/internal/service"

	commentV1 "kratos-cms/gen/api/go/comment/service/v1"
	contentV1 "kratos-cms/gen/api/go/content/service/v1"
	fileV1 "kratos-cms/gen/api/go/file/service/v1"
	userV1 "kratos-cms/gen/api/go/user/service/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	commentService *service.CommentService,

	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	tagSvc *service.TagService,

	userSvc *service.UserService,

	attachmentSvc *service.AttachmentService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))

	commentV1.RegisterCommentServiceServer(srv, commentService)

	contentV1.RegisterPostServiceServer(srv, postSvc)
	contentV1.RegisterLinkServiceServer(srv, linkSvc)
	contentV1.RegisterCategoryServiceServer(srv, cateSvc)
	contentV1.RegisterTagServiceServer(srv, tagSvc)

	fileV1.RegisterAttachmentServiceServer(srv, attachmentSvc)

	userV1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
