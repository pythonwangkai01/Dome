package model

import (
	"time"
	"user/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
}
