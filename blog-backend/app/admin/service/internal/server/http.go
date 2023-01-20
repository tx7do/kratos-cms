package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"

	"github.com/gorilla/handlers"

	"github.com/tx7do/kratos-authn/authn"
	"github.com/tx7do/kratos-authn/engine/jwt"

	"github.com/tx7do/kratos-authz/authz"
	"github.com/tx7do/kratos-authz/engine/noop"

	v1 "kratos-blog/gen/api/go/admin/service/v1"
	"kratos-blog/pkg/middleware/auth"

	"kratos-blog/app/admin/service/internal/conf"
	"kratos-blog/app/admin/service/internal/service"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[v1.OperationUserServiceRegister] = true
	whiteList[v1.OperationUserServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func NewMiddleware(ac *conf.Auth, logger log.Logger) http.ServerOption {
	authenticator, _ := jwt.NewAuthenticator(ac.ApiKey, "HS256")
	authorizer := noop.State{}

	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			authn.Server(authenticator),
			auth.Server(),
			authz.Server(authorizer),
		).Match(NewWhiteListMatcher()).Build(),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server, ac *conf.Auth, logger log.Logger,
	userSvc *service.UserService,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	commentSvc *service.CommentService,
	tagSvc *service.TagService,
	attachmentSvc *service.AttachmentService,
) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(ac, logger),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	srv.HandlePrefix("/q/", openapiv2.NewHandler())

	v1.RegisterPostServiceHTTPServer(srv, postSvc)
	v1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	v1.RegisterTagServiceHTTPServer(srv, tagSvc)
	v1.RegisterLinkServiceHTTPServer(srv, linkSvc)
	v1.RegisterUserServiceHTTPServer(srv, userSvc)
	v1.RegisterAttachmentServiceHTTPServer(srv, attachmentSvc)
	v1.RegisterCommentServiceHTTPServer(srv, commentSvc)

	return srv
}
