package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App       *App
	Log       *LogConfig
	MySQLConf *DBConfig
}

type App struct {
	HttpPort     int
	RunMode      string
	ReadTimeout  int //单位s
	WriteTimeout int
}

func New() *Config {
	ReadConfFromFile()

	return &Config{
		App:       GetAppConfig(),
		Log:       GetLogConfig(),
		MySQLConf: GetDataSourceConfig(),
	}
}

// ReadConfFromFile viper初始化完成，就可以直接使用viper.get来参与运算了
//当需要读多个文件的时候，需要用户管理多个viper实例，不能默认使用它的单例
func ReadConfFromFile() {
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

func GetAppConfig() *App {
	return &App{
		HttpPort:     viper.GetInt("http.port"),
		RunMode:      viper.GetString("http.mode"),
		ReadTimeout:  viper.GetInt("http.read-time-out"),
		WriteTimeout: viper.GetInt("http.write-time-out"),
	}
}

func GetLogConfig() *LogConfig {
	return &LogConfig{
		Level:      viper.GetString("log.level"),
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxAge:     viper.GetInt("log.maxAge"),
		MaxBackups: viper.GetInt("log.maxBackups"),
	}
}

func GetDataSourceConfig() *DBConfig {
	return &DBConfig{
		Addr:         viper.GetString("datasource.addr"),
		DSN:          viper.GetString("datasource.dsn"),
		ReadDSN:      viper.GetStringSlice("datasource.read-dsn"),
		Active:       viper.GetInt("datasource.active"),
		Idle:         viper.GetInt("datasource.idle"),
		IdleTimeout:  viper.GetInt("datasource.idle-time-out"),
		QueryTimeout: viper.GetInt("datasource.query-time-out"),
		ExecTimeout:  viper.GetInt("datasource.exec-time-out"),
		TranTimeout:  viper.GetInt("datasource.tran-time-out"),
	}

}
