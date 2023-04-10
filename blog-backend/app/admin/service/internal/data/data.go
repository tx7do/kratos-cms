package data

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	commentV1 "kratos-blog/gen/api/go/comment/service/v1"
	contentV1 "kratos-blog/gen/api/go/content/service/v1"
	fileV1 "kratos-blog/gen/api/go/file/service/v1"
	userV1 "kratos-blog/gen/api/go/user/service/v1"

	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/bootstrap"
	"kratos-blog/pkg/service"
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine

	userClient       userV1.UserServiceClient
	attachmentClient fileV1.AttachmentServiceClient
	commentClient    commentV1.CommentServiceClient
	categoryClient   contentV1.CategoryServiceClient
	linkClient       contentV1.LinkServiceClient
	postClient       contentV1.PostServiceClient
	tagClient        contentV1.TagServiceClient
}

// NewData .
func NewData(redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	userClient userV1.UserServiceClient,
	attachmentClient fileV1.AttachmentServiceClient,
	commentClient commentV1.CommentServiceClient,
	categoryClient contentV1.CategoryServiceClient,
	linkClient contentV1.LinkServiceClient,
	postClient contentV1.PostServiceClient,
	tagClient contentV1.TagServiceClient,
	logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		rdb:              redisClient,
		log:              l,
		authenticator:    authenticator,
		authorizer:       authorizer,
		userClient:       userClient,
		attachmentClient: attachmentClient,
		commentClient:    commentClient,
		categoryClient:   categoryClient,
		linkClient:       linkClient,
		postClient:       postClient,
		tagClient:        tagClient,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Data.Redis.Addr,
		Password:     cfg.Data.Redis.Password,
		DB:           int(cfg.Data.Redis.Db),
		DialTimeout:  cfg.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: cfg.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.Data.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		l.Fatalf("failed opening connection to redis")
	}
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Auth.ApiKey)),
		jwt.WithSigningMethod("HS256"),
	)
	return authenticator
}

// NewAuthorizer 创建权鉴器
func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewAttachmentServiceClient(r registry.Discovery, c *conf.Bootstrap) fileV1.AttachmentServiceClient {
	return fileV1.NewAttachmentServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewCommentServiceClient(r registry.Discovery, c *conf.Bootstrap) commentV1.CommentServiceClient {
	return commentV1.NewCommentServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewCategoryServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.CategoryServiceClient {
	return contentV1.NewCategoryServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewLinkServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.LinkServiceClient {
	return contentV1.NewLinkServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewPostServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.PostServiceClient {
	return contentV1.NewPostServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}

func NewTagServiceClient(r registry.Discovery, c *conf.Bootstrap) contentV1.TagServiceClient {
	return contentV1.NewTagServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, c.Server.Grpc.GetTimeout()))
}
