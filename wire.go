//go:build wireinject

package main

import (
	"chat/biz/repository"
	"chat/config"

	"github.com/google/wire"
)

var InitConfigSet = wire.NewSet(config.InitConfig, config.ConfigPath)

// 初始化所有关于数据库的依赖
func Initialize() error {
	wire.Build(initDB, InitConfigSet)
	return nil
}
func initDB(c config.Config) error {
	wire.Build(repository.InitRepo, repository.InitDB)
	return nil
}
