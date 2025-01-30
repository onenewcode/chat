package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	// 创建一个模拟的 Viper 对象
	vp := viper.New()

	// 设置一些模拟的配置数据
	vp.Set("redis.addr", "localhost:6379")

	// 调用待测试的函数
	config := InitConfig(vp)

	// 验证解析后的配置是否正确
	if config.Redis.Addr != "localhost:6379" {
		t.Errorf("解析后的 Redis 地址不正确，期望: localhost:6379, 实际: %s", config.Redis.Addr)
	}
}
