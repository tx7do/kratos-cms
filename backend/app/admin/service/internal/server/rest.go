package server

import (
	"context"
	"kratos-cms/pkg/cache"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-cms/app/admin/service/cmd/server/assets"
	"kratos-cms/app/admin/service/internal/service"

	adminV1 "kratos-cms/api/gen/go/admin/service/v1"
	"kratos-cms/pkg/middleware/auth"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[adminV1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(
	logger log.Logger,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	userToken *cache.UserToken,
) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(userToken),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	userToken *cache.UserToken,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
	postSvc *service.PostService,
	linkSvc *service.LinkService,
	cateSvc *service.CategoryService,
	commentSvc *service.CommentService,
	tagSvc *service.TagService,
	attachmentSvc *service.AttachmentService,
) *http.Server {
	srv := rpc.CreateRestServer(cfg, newRestMiddleware(logger, authenticator, authorizer, userToken)...)

	adminV1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	adminV1.RegisterPostServiceHTTPServer(srv, postSvc)
	adminV1.RegisterCategoryServiceHTTPServer(srv, cateSvc)
	adminV1.RegisterTagServiceHTTPServer(srv, tagSvc)
	adminV1.RegisterLinkServiceHTTPServer(srv, linkSvc)
	adminV1.RegisterUserServiceHTTPServer(srv, userSvc)
	adminV1.RegisterAttachmentServiceHTTPServer(srv, attachmentSvc)
	adminV1.RegisterCommentServiceHTTPServer(srv, commentSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Admin Service"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
