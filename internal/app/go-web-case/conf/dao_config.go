package conf

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DataSourceConf struct {
	MaxConn int //最大连接数
	MaxOpen int
	Dsn     string
}

var globalDB *gorm.DB

func InitDB(cfg *DataSourceConf) (err error) {
	zap.L().Info("db init  ", zap.String("dsn", cfg.Dsn))
	ormLogger := logger.Default

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表名前缀
		},
	})

	//根据 *gorm.DB 获取 *sql.DB 来配置线程池
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		zap.L().Fatal(err.Error())
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpen)

	globalDB = db

	zap.L().Info("db connected success")
	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}
