package test

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web-case/config"
	"testing"
)

func TestInitConf(t *testing.T) {
	config.InitConf()
	fmt.Println(viper.Get("bigbro.diaodu_url"))
	fmt.Println(viper.Get("auth.server.url"))
}

func TestEnv(t *testing.T) {

}
