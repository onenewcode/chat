package redis

import (
	"chat/utils"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

/*
*
设置在线用户到redis缓存
*
*/
func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	utils.Red.Set(ctx, key, val, timeTTL)
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	hlog.Info("Publish 。。。。", msg)
	err = redisClient.Publish(ctx, channel, msg).Err()
	if err != nil {
		hlog.Error(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := redisClient.Subscribe(ctx, channel)
	hlog.Info("Subscribe 。。。。", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		hlog.Info(err)
		return "", err
	}
	hlog.Info("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
