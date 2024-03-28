package starter

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"simple-skeleton/boot"
)

var cache *redis.Client

func Cache() *redis.Client {
	return cache
}

type RedisStater struct {
	boot.BaseStarter
}

func (r *RedisStater) Setup(ctx boot.StaterContext) {
	conf := ctx.Conf()
	cache = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Redis.Addr, conf.Redis.Port),
		Password:     conf.Redis.Password,
		DB:           conf.Redis.DB,
		DialTimeout:  conf.Redis.DialTimeout,
		ReadTimeout:  conf.Redis.ReadTimeout,
		WriteTimeout: conf.Redis.WriteTimeout,
		PoolSize:     conf.Redis.PoolSize,
		MinIdleConns: conf.Redis.MinIdleConns,
	})

	err := cache.Ping().Err()
	if err != nil {
		panic(err)
	}

	logrus.Info("redis connect success!!!")
}
