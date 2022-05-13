package conf

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	Addr         string   // for trace
	DSN          string   // write data source name.
	ReadDSN      []string // read data source name.
	Active       int      // pool
	Idle         int      // pool
	IdleTimeout  int      // connect max life time.
	QueryTimeout int      // query sql timeout
	ExecTimeout  int      // execute sql timeout
	TranTimeout  int      // transaction sql timeout
}

func InitDB(cfg *DBConfig) *gorm.DB {
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
		return nil
	}

	sqlDB.SetMaxIdleConns(cfg.Idle)
	sqlDB.SetMaxOpenConns(cfg.Active)

	zap.L().Info("db connected success")
	return db
}

