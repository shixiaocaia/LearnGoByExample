package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_example/web_app/settings"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
