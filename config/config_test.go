package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	vp := viper.New()
	vp.SetConfigName("app")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yml")
	config := InitConfig(vp)
	println(config.Redis.Addr)
}
