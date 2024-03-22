package config

import (
	"github.com/spf13/viper"
)

// 全局变量，提供给内部的其他包使用
var (
	Redis   redis
	Mysql   mysql
	Timeout timeout
	Port    port
)

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
	HeartbeatHz      int
	HeartbeatMaxTime int
	RedisOnlineTime  int
}
type port struct {
	Server string
	Udp    int
}

// 初始化一个配置类，让viper读取指定的配置文件
func configPath() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("app")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return vp, nil
}

func readSection(vp *viper.Viper, k string, v interface{}) error {
	err := vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

// 初始化配置，把所有的数据读取后放入global的全局变量中
func InitConfig() {
	vp, err := configPath()
	if err != nil {
		panic("配置文件读取错误")
	}
	err = readSection(vp, "redis", &Redis)
	if err != nil {
		panic("redis类读取错误，检查server类映射是否正确")
	}
	err = readSection(vp, "mysql", &Mysql)
	if err != nil {
		panic("mysql类读取错误，检查App类映射是否正确")
	}
	err = readSection(vp, "timeout", &Timeout)
	if err != nil {
		panic("timeout类读取错误，检查Database类映射是否正确")
	}
	err = readSection(vp, "port", &Port)
	if err != nil {
		panic("port类读取错误，检查port类映射是否正确")
	}
}
