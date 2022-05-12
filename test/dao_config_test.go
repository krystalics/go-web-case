package test

import (
	"go-web-case/internal/app/go-web-case/conf"
	"testing"
)

func TestName(t *testing.T) {
	conf.InitDB(conf.GetDataSourceConfig())
}
