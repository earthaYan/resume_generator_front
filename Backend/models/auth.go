package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12 //密码加密难度
)

type User struct {
	gorm.Model
	UserName        string `gorm:"unique"`
	PasswordDigiest string
}
type UserServiceRegisterReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigiest = string(bytes)
	return nil
}

type UserServiceLoginReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (u *User) CheckPwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigiest), []byte(password))
	return err == nil
}

type UserLoginResponse struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	CreatedAt int64  `jsonL:"created_at"`
}
type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}
