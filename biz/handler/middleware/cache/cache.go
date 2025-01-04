package cache

import (
	"chat/common"
	"context"
)

// 删除所有用户
func DeleteALLFriends() {
	redisStore.Delete(context.Background(), common.ALLFriends)
	redisStore.Delete(context.Background(), common.ALLUserFriends)
}
