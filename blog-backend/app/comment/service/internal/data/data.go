package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"kratos-blog/app/comment/service/internal/data/ent"
	"kratos-blog/app/comment/service/internal/data/ent/migrate"
	"kratos-blog/gen/api/go/common/conf"
)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(entClient *ent.Client, redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/comment-service"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/comment-service"))

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

// NewEntClient 创建数据库客户端
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	l := log.NewHelper(log.With(logger, "module", "ent/data/comment-service"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	// 运行数据库迁移工具
	if false {
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client
}
