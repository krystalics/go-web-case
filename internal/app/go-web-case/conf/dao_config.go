package conf

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type DBConfig struct {
	Addr         string        // for trace
	DSN          string        // write data source name.
	ReadDSN      []string      // read data source name.
	Active       int           // pool
	Idle         int           // pool
	IdleTimeout  time.Duration // connect max life time.
	QueryTimeout time.Duration // query sql timeout
	ExecTimeout  time.Duration // execute sql timeout
	TranTimeout  time.Duration // transaction sql timeout
}

var globalDB *gorm.DB

func InitDB(cfg *DBConfig) (err error) {
	zap.L().Info("db init  ", zap.String("dsn", cfg.DSN))
	ormLogger := logger.Default

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
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

	sqlDB.SetMaxIdleConns(cfg.Idle)
	sqlDB.SetMaxOpenConns(cfg.Active)

	globalDB = db

	zap.L().Info("db connected success")
	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}
