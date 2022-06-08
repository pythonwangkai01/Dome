package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid            uint   `gorm:"not null";gorm:"unique"`
	Phone          string `gorm:"unique"`
	UserName       string `gorm:"unique"`
	Address        string
	PasswordDigest string
	Desc           string
	Status         int
}

type AdminUser struct {
	gorm.Model
	Uid            uint   `gorm:"not null";gorm:"unique"`
	Phone          string `gorm:"unique"`
	UserName       string `gorm:"unique"`
	PasswordDigest string
	CreateUid      uint `gorm:"not null"`
}

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{}, &AdminUser{})
}
