package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	registerGorm()
}

func registerGorm() {
	dsn := "root:123456@tcp(127.0.0.1:36000)/user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: loggerConfig})
	if err != nil {
		panic(err)
	}
}
