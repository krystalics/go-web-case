package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//viper初始化完成，就可以直接使用viper.get来参与运算了
//当需要读多个文件的时候，需要用户管理多个viper实例，不能默认使用它的单例
func InitConf() {
	suffix := os.Getenv("APP_ENV")
	if suffix != "" {
		viper.SetConfigName("application" + "-" + suffix)
	} else {
		viper.SetConfigName("application")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs") //和入口main文件的相对位置
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
