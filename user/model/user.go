package model

import (
	"user/pkg/e"

	"golang.org/x/crypto/bcrypt"
)

const PassWordCose = 9

//加密密码
func (user *User) SetPassWord(password string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCose)
	e.HandlerError(err, `bcrypt.GenerateFromPassword`)
	user.PasswordDigest = string(b)
	return nil
}

//检验密码
func (user *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

//前端md5密码解密
func (user *User) Md5PassWord(passwword string) string {

}
