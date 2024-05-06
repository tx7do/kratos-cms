package data

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/redis/go-redis/v9"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/tx7do/go-utils/entgo"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"

	"kratos-cms/app/core/service/internal/data/ent"
)

// Data .
type Data struct {
	log *log.Helper
	db  *entgo.EntClient[*ent.Client]
	rdb *redis.Client
}

// NewData .
func NewData(entClient *entgo.EntClient[*ent.Client], redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/core-service"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		_ = d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/core-service"))
	return redisClient.NewClient(cfg.Data)
}
