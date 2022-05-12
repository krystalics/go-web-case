package test

import (
	"go-web-case/internal/app/go-web-case/conf"
	"gorm.io/driver/mysql"
	"testing"
)

func TestName(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/permission?charset=utf8mb4&parseTime=True&loc=Local"

	conf.Connect(&mysql.Config{
		DSN:                       dsn,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	})

}
