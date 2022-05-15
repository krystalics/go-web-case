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
	HttpPort     int    `mapstructure:"port"`
	RunMode      string `mapstructure:"mode"`
	ReadTimeout  int    `mapstructure:"read-time-out"`
	WriteTimeout int    `mapstructure:"write-time-out"`
}

func New() *Config {
	ReadConfFromFile()

	return &Config{
		App:       NewSubConfig[App]("http"),
		Log:       NewSubConfig[LogConfig]("log"),
		MySQLConf: NewSubConfig[DBConfig]("datasource"),
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

// NewSubConfig 泛型 真爽
func NewSubConfig[T any](key string) *T {
	var t T
	err := viper.Sub(key).Unmarshal(&t)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return &t
}
