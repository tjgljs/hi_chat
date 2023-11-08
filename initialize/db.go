package initialize

import (
	"HiChat/global"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {

	//dsn:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",User,Password,Host,Port,DBName)
	dsn := "root:123456789@tcp(localhost:3306)/hi_chat?charset=utf8&parseTime=True&loc=Local"
	//sql日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, //慢sql阈值
			LogLevel:                  logger.Info, //日志级别
			IgnoreRecordNotFoundError: true,        //忽略未找到的错误
			Colorful:                  true,        //禁止彩色打印
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func InitRedis() {
	opt := redis.Options{
		// Addr:     fmt.Sprintf("%s:%d", global.ServiceConfig.RedisDB.Host, global.ServiceConfig.RedisDB.Port),
		// redis地址
		Addr:     "localhost:6379",
		Password: "", // redis密码，没有则留空
		DB:       10, // 默认数据库，默认是0
	}
	global.RedisDB = redis.NewClient(&opt)
}
