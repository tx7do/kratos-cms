package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-cms/app/core/service/internal/service"

	commentV1 "kratos-cms/api/gen/go/comment/service/v1"
	contentV1 "kratos-cms/api/gen/go/content/service/v1"
	fileV1 "kratos-cms/api/gen/go/file/service/v1"
	userV1 "kratos-cms/api/gen/go/user/service/v1"
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
	fileSvc *service.FileService,
) *grpc.Server {
	srv := rpc.CreateGrpcServer(cfg, logging.Server(logger))

	commentV1.RegisterCommentServiceServer(srv, commentService)

	contentV1.RegisterPostServiceServer(srv, postSvc)
	contentV1.RegisterLinkServiceServer(srv, linkSvc)
	contentV1.RegisterCategoryServiceServer(srv, cateSvc)
	contentV1.RegisterTagServiceServer(srv, tagSvc)

	fileV1.RegisterAttachmentServiceServer(srv, attachmentSvc)
	fileV1.RegisterFileServiceServer(srv, fileSvc)

	userV1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
