package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	"kratos-blog/app/front/service/internal/service"
	"kratos-blog/gen/api/go/common/conf"
	v1 "kratos-blog/gen/api/go/front/service/v1"
	"kratos-blog/pkg/bootstrap"
	"kratos-blog/pkg/middleware/auth"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(authenticator authnEngine.Authenticator, authorizer authzEngine.Engine, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	commentSvc *service.CommentService,
	tagSvc *service.TagService,
	attachmentSvc *service.AttachmentService,
) *http.Server {
	srv := bootstrap.CreateRestServer(cfg, newRestMiddleware(authenticator, authorizer, logger)...)

	v1.RegisterPostServiceHTTPServer(srv, postSvc)
	v1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	v1.RegisterTagServiceHTTPServer(srv, tagSvc)
	v1.RegisterLinkServiceHTTPServer(srv, linkSvc)
	v1.RegisterAttachmentServiceHTTPServer(srv, attachmentSvc)
	v1.RegisterCommentServiceHTTPServer(srv, commentSvc)

	return srv
}
