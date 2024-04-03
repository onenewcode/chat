package utils

import (
	"chat/common"
	"context"
	"github.com/hertz-contrib/cache/persist"
)

var (
	// 缓存列表
	RedisStore *persist.RedisStore
)

// 删除所有用户
func DeleteALLFriends() {
	RedisStore.Delete(context.Background(), common.ALLFriends)
	RedisStore.Delete(context.Background(), common.ALLUserFriends)
}
