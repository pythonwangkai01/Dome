package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid            uint `gorm:"not null";`
	Phone          string
	Email          string
	UserName       string `gorm:"unique"`
	PasswordDigest string
	Desc           string
	Status         int
}

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
}
