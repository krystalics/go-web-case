package config

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

type Dao struct {
}

var globalDB *gorm.DB

//在项目初始化前、调用它
func Connect(cfg *mysql.Config) {

	dsn := fmt.Sprintf(
		"%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DSN,
	)
	globalLogger.Debugf("db dsn: %s", dsn)

	ormLogger := logger.Default

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表名前缀
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	globalDB = db

	globalLogger.Errorf("db connected success")
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}