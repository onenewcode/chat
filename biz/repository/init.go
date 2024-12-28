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
	"gorm.io/plugin/prometheus"
)

var DB *gorm.DB

// 初始化数据库库
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

// initPrometheus
// Note: https://gorm.io/zh_CN/docs/prometheus.html
func initPrometheus(db *gorm.DB) {
	db.Use(prometheus.New(prometheus.Config{
		DBName:          "db1",                       // 使用 `DBName` 作为指标 label
		RefreshInterval: 15,                          // 指标刷新频率（默认为 15 秒）
		PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
		StartServer:     true,                        // 启用一个 http 服务来暴露指标
		HTTPServerPort:  8080,                        // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // 用户自定义指标
	}))
}
