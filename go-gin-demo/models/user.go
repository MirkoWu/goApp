package models

import (
	"fmt"
	"time"
)

type User struct {
	UserId        int64  `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Age           int    `json:"age"`
	Signature     string `json:"signature"`
	RegisterTime  int64  `json:"register_time"`
	LastLoginTime int64  `json:"last_login_time"`
}

func GetUserByID(userId int64) (user User) {
	db.Select("user").Where("user_id = ?", userId).First(&user)
	return
}

func GetUserByEmail(email string) (user User) {
	db.Where("email = ?", email).First(&user)
	fmt.Println(user.Password)
	return
}

func GetUserTotal() (count int) {
	db.Model(&User{}).Count(&count)
	return
}

func ExistUserByEmail(email string) bool {
	var user User
	db.Where("email = ?", email).First(&user)
	if user.UserId > 0 {
		return true
	}

	return false
}
func ExistUserByID(userId int64) bool {
	var user User
	db.Select("user").Where("id = ?", userId).First(&user)
	if user.UserId > 0 {
		return true
	}

	return false
}

//func DeleteUser(id int  ) bool {
//	db. Where("id = ?", id).Delete(&Tag{})
//
//	return true
//}

func EditUser(id int64, data interface{}) bool {
	db.Model(&User{}).Where("id = ?", id).Update(data)

	return true
}

func AddUser(email string, password string) bool {
	//TimeFormat := "20060102 12:12:12"
	db.Create(&User{
		UserId:        1,
		Email:         email,
		Password:      password,
		Nickname:      "昵称",
		Avatar:        "s",
		Age:           -1,
		Signature:     "...",
		RegisterTime:  time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	})

	return true
}
