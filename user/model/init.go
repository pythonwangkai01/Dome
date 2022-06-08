package model

import (
	"time"
	"user/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func DataBase(connString string) {
	db, err := gorm.Open("mysql", connString)
	e.HandlerError(err, `gorm.Open("mysql", connString)`)

	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)
	//一些配置
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 25)

	DB = db
	migration()
	//init admin
	initAdmin()
}

func SetPassWord(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCose)
	e.HandlerError(err, `bcrypt.GenerateFromPassword`)
	PasswordDigest := string(b)
	return PasswordDigest
}

func initAdmin() {
	adminuser := AdminUser{
		UserName:       "admin",
		PasswordDigest: SetPassWord("123456"),
		Phone:          "18717908454",
		Uid:            1314521,
		CreateUid:      1314521,
	}
	err := DB.Where("user_name=?", adminuser.UserName).First(&adminuser).Error
	if err != nil {
		//未找到进行
		DB.Create(&adminuser).Update("CreatedAt", time.Now())
	}

}
