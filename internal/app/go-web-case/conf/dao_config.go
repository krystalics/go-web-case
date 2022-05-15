package conf

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DBConfig struct {
	Addr         string   `mapstructure:"addr"`          // for trace
	DSN          string   `mapstructure:"dsn"`           // write data source name.
	ReadDSN      []string `mapstructure:"read-dsn"`      // read data source name.
	Active       int      `mapstructure:"active"`        // pool
	Idle         int      `mapstructure:"idle"`          // pool
	IdleTimeout  int      `mapstructure:"idle-timeout"`  // connect max life time.
	QueryTimeout int      `mapstructure:"query-timeout"` // query sql timeout
	ExecTimeout  int      `mapstructure:"exec-timeout"`  // execute sql timeout
	TranTimeout  int      `mapstructure:"tran-timeout"`  // transaction sql timeout
}

func InitDB(cfg *Config) *gorm.DB {
	//todo 临时采用新建writer的方式、将日志转而写入app.log中、后续再调查其他方案
	var l logger.Interface
	if cfg.App.RunMode == "debug" {
		l = logger.Default
	} else {
		writer := getLogWriter(cfg.Log.Filename, cfg.Log.MaxSize, cfg.Log.MaxBackups, cfg.Log.MaxAge)
		l = logger.New(log.New(writer, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		})
	}

	db, err := gorm.Open(mysql.Open(cfg.MySQLConf.DSN), &gorm.Config{
		Logger: l,
	})

	//根据 *gorm.DB 获取 *sql.DB 来配置线程池
	sqlDB, err := db.DB()
	//defer sqlDB.Close()

	if err != nil {
		zap.L().Fatal(err.Error())
		return nil
	}

	sqlDB.SetMaxIdleConns(cfg.MySQLConf.Idle)
	sqlDB.SetMaxOpenConns(cfg.MySQLConf.Active)
	sqlDB.Ping()

	//sqlDB.SetConnMaxLifetime(cfg.)
	zap.L().Info("db connected success")
	return db
}
