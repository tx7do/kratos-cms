package data

import (
	"context"
	"time"

	GRPC "google.golang.org/grpc"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	commentV1 "kratos-blog/gen/api/go/comment/service/v1"
	contentV1 "kratos-blog/gen/api/go/content/service/v1"
	fileV1 "kratos-blog/gen/api/go/file/service/v1"
	userV1 "kratos-blog/gen/api/go/user/service/v1"

	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/service"
	"kratos-blog/pkg/util/bootstrap"
)

const defaultTimeout = 5 * time.Second

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

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
func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		l.Fatalf("failed opening connection to redis")
	}
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	return bootstrap.NewConsulRegistry(conf.Consul.Address, conf.Consul.Scheme, conf.Consul.HealthCheck)
}

func createGrpcConnection(serviceName string, r registry.Discovery, c *conf.Server) GRPC.ClientConnInterface {
	timeout := defaultTimeout
	if c.Grpc != nil && c.Grpc.Timeout != nil {
		timeout = c.Grpc.GetTimeout().AsDuration()
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(timeout),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}
	return conn
}

func NewUserServiceClient(r registry.Discovery, c *conf.Server) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(createGrpcConnection(service.UserService, r, c))
}

func NewAttachmentServiceClient(r registry.Discovery, c *conf.Server) fileV1.AttachmentServiceClient {
	return fileV1.NewAttachmentServiceClient(createGrpcConnection(service.FileService, r, c))
}

func NewCommentServiceClient(r registry.Discovery, c *conf.Server) commentV1.CommentServiceClient {
	return commentV1.NewCommentServiceClient(createGrpcConnection(service.CommentService, r, c))
}

func NewCategoryServiceClient(r registry.Discovery, c *conf.Server) contentV1.CategoryServiceClient {
	return contentV1.NewCategoryServiceClient(createGrpcConnection(service.ContentService, r, c))
}

func NewLinkServiceClient(r registry.Discovery, c *conf.Server) contentV1.LinkServiceClient {
	return contentV1.NewLinkServiceClient(createGrpcConnection(service.ContentService, r, c))
}

func NewPostServiceClient(r registry.Discovery, c *conf.Server) contentV1.PostServiceClient {
	return contentV1.NewPostServiceClient(createGrpcConnection(service.ContentService, r, c))
}

func NewTagServiceClient(r registry.Discovery, c *conf.Server) contentV1.TagServiceClient {
	return contentV1.NewTagServiceClient(createGrpcConnection(service.ContentService, r, c))
}
