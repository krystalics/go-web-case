package test

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web-case/internal/app/go-web-case/conf"
	"testing"
)

func TestInitConf(t *testing.T) {
	conf.InitConf()
	fmt.Println(viper.Get("bigbro.diaodu_url"))
	fmt.Println(viper.Get("auth.server.url"))
}

func TestEnv(t *testing.T) {

}
