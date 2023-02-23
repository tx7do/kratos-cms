package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	"kratos-blog/gen/api/go/common/conf"
)

// NewRedisClient 创建Redis客户端
func NewRedisClient(conf *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		log.Fatalf("failed opening connection to redis")
		return nil
	}
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}
