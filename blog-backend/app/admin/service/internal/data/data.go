package data

import (
	"context"
	"time"

	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	commentV1 "kratos-blog/gen/api/go/comment/service/v1"
	contentV1 "kratos-blog/gen/api/go/content/service/v1"
	fileV1 "kratos-blog/gen/api/go/file/service/v1"
	userV1 "kratos-blog/gen/api/go/user/service/v1"

	"kratos-blog/app/admin/service/internal/conf"
	"kratos-blog/pkg/service"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewDiscovery,

	NewUserServiceClient,
	NewAttachmentServiceClient,
	NewCommentServiceClient,
	NewCategoryServiceClient,
	NewLinkServiceClient,
	NewPostServiceClient,
	NewTagServiceClient,
)

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
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(conf.Consul.HealthCheck))
	return r
}

func NewUserServiceClient(r registry.Discovery) userV1.UserServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.UserService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserServiceClient(conn)
}

func NewAttachmentServiceClient(r registry.Discovery) fileV1.AttachmentServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.FileService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return fileV1.NewAttachmentServiceClient(conn)
}

func NewCommentServiceClient(r registry.Discovery) commentV1.CommentServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.CommentService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return commentV1.NewCommentServiceClient(conn)
}

func NewCategoryServiceClient(r registry.Discovery) contentV1.CategoryServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.ContentService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return contentV1.NewCategoryServiceClient(conn)
}

func NewLinkServiceClient(r registry.Discovery) contentV1.LinkServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.ContentService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return contentV1.NewLinkServiceClient(conn)
}

func NewPostServiceClient(r registry.Discovery) contentV1.PostServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.ContentService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return contentV1.NewPostServiceClient(conn)
}

func NewTagServiceClient(r registry.Discovery) contentV1.TagServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.ContentService),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return contentV1.NewTagServiceClient(conn)
}
