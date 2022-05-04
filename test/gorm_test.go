package test

import (
	"go-web-case/internal/app/go-web-case/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGORM(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/permission?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect db")
	}

	//db.AutoMigrate(&model.User{})
	db.Create(&model.User{
		Ucid:      1,
		UserName:  "林家宝",
		Email:     "linjiabao001@ss.com",
		IsDeleted: false,
	})

	var user model.User
	db.First(&user, 1)
	db.First(&user, "Email=?", "linjiabao001@ss.com")

	db.Model(&user).Where("ucid=?", 1).Update("UserName", "linjiabao")
	db.Model(&user).Where("ucid=?", 1).Updates(model.User{
		UserName: "linjiabao2",
		Email:    "email",
	})
	db.Model(&user).Where("ucid=?", 1).Updates(map[string]interface{}{
		"UserName": "mapt",
		"Email":    "maptest",
	})

	db.Delete(&user, 1)

}
