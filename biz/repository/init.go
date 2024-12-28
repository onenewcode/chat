package repository

import (
	"chat/config"
	"log"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(c config.Config) {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:      logger.Info, // 设置日志级别
			SlowThreshold: time.Second, // 慢查询阈值
			Colorful:      true,        // 是否使用彩色日志
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DNS),
		&gorm.Config{Logger: dbLogger})
	if err != nil {
		hlog.Error(err)
		panic("failed to connect database")
	}
	hlog.Info(" MySQL init 。。。。")
}
