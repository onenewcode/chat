package redis

import (
	"chat/config"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

const (
	PublishKey = "websocket"
)

func InitRedis(c config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           c.Redis.DB,
		PoolSize:     c.Redis.PoolSize,
		MinIdleConns: c.Redis.MinIdleConn,
	})
}
