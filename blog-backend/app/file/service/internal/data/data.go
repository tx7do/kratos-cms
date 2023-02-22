package data

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	"kratos-blog/app/file/service/internal/data/model"
	"kratos-blog/gen/api/go/common/conf"
)

// Data .
type Data struct {
	log *log.Helper
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(gormClient *gorm.DB, redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/file-service"))

	d := &Data{
		db:  gormClient,
		rdb: redisClient,
		log: l,
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
	l := log.NewHelper(log.With(logger, "module", "redis/data/file-service"))

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

func createGormDialector(driver string, dsn string) gorm.Dialector {
	switch driver {
	case "mysql":
		return mysql.Open(dsn)
	case "postgres":
		return postgres.Open(dsn)
	case "sqlserver":
		return sqlserver.Open(dsn)
	case "sqlite":
		return sqlite.Open(dsn)
	case "clickhouse":
		return clickhouse.Open(dsn)
	default:
		return nil
	}
}

// NewGormClient 创建数据库客户端
func NewGormClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "db/data/file-service"))

	var client *gorm.DB
	var err error

	// 连接数据库
	client, err = gorm.Open(createGormDialector(conf.Database.Driver, conf.Database.Source), &gorm.Config{})

	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	// 运行数据库迁移工具
	if true {
		if err := client.AutoMigrate(&model.Attachment{}); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return client
}
