//go:build wireinject

package main

import (
	"chat/config"

	"github.com/google/wire"
)

func Initialize() config.Config {
	wire.Build(config.InitConfig, config.ConfigPath)
	return config.Config{}
}
