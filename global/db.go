package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Object.MysqlUsername,
		Object.MysqlPassword,
		Object.MysqlHost,
		Object.MysqlPort,
		Object.MysqlDbname)
	globalDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	sqlDB, err := globalDB.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(12 * time.Hour)

	Log.Infof("connect to mysql %s:%d, dbname=%s successfully",
		Object.MysqlHost,
		Object.MysqlPort,
		Object.MysqlDbname)

	return globalDB
}
