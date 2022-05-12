package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web-case/internal/app/go-web-case/conf"
	"os"
)

// InitConf viper初始化完成，就可以直接使用viper.get来参与运算了
//当需要读多个文件的时候，需要用户管理多个viper实例，不能默认使用它的单例
func InitConf() {
	suffix := os.Getenv("APP_ENV")
	if suffix != "" {
		viper.SetConfigName("application" + "-" + suffix)
	} else {
		viper.SetConfigName("application")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") //和入口main文件的相对位置
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func GetLogConfig() *conf.LogConfig {
	return &conf.LogConfig{
		Level:      viper.GetString("log.level"),
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxAge:     viper.GetInt("log.maxAge"),
		MaxBackups: viper.GetInt("log.maxBackups"),
	}
}

func GetDataSourceConfig() *conf.DataSourceConf {
	return &conf.DataSourceConf{
		MaxOpen: viper.GetInt("datasource.maxOpen"),
		MaxConn: viper.GetInt("datasource.maxConn"),
		Dsn:     viper.GetString("datasource.dsn"),
	}
}