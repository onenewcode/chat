package cache

import (
	"github.com/hertz-contrib/cache/persist"
)

var (
	// 缓存列表
	redisStore *persist.RedisStore
)
