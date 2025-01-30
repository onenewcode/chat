package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 全局变量，提供给内部的其他包使用
var (
	Redis        redis
	Mysql        mysql
	Timeout      timeout
	Port         port
	GlobalConfig Config
)

type Config struct {
	Redis   redis
	Mysql   mysql
	Timeout timeout
	Port    port
}
type mysql struct {
	DNS string
}
type redis struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	MinIdleConn int
}
type timeout struct {
	DelayHeartbeat   int
	HeartbeatHz      uint64
	HeartbeatMaxTime int
	RedisOnlineTime  int
}
type port struct {
	Server string
	Udp    int
}

// 初始化一个配置类，让 viper 读取指定的配置文件
func ConfigPath() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("app")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yml")
	err := vp.ReadInConfig()
	if err != nil {
		panic("配置文件读取错误")
	}

	return vp
}

// 初始化配置，把所有的数据读取后放入 global 的全局变量中
func InitConfig(vp *viper.Viper) Config {
	var config Config
	if err := vp.Unmarshal(&config); err != nil {
		fmt.Printf("解析字段失败, %v", err)
	}
	return config
}
