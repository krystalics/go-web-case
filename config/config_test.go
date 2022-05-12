package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestInitConf(t *testing.T) {
	InitConf()
	fmt.Println(viper.Get("bigbro.diaodu_url"))
	fmt.Println(viper.Get("auth.server.url"))
}

func TestEnv(t *testing.T) {

}
