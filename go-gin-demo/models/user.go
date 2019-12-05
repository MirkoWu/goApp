package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"time"
)

type User struct {
	Model
	UserId        int    `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"-"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Sex           int    `json:"sex"`
	Signature     string `json:"signature"`
	RegisterTime  int64  `json:"register_time"`
	LastLoginTime int64  `json:"last_login_time"`
	Token         string `json:"token"`
}

func GetUserByID(userId int) (user User) {
	db.Where("user_id = ?", userId).First(&user)
	return
}

func GetUserByEmail(email string) (user User) {
	db.Where("email = ?", email).First(&user)
	return
}

//func GetUserWithPwdByEmail(email string) (user User) {
//	db.Where("email = ?", email).First(&user)
//	fmt.Println(user.Password)
//	return
//}

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

//func ExistUserByID(userId int) (isExist bool) {
//	if userId <= 0 {
//		return false
//	}
//var user User
//	db.Where("user_id = ?", userId).First(&user)
//	if user.UserId > 0 {
//		return true
//	}
//
//	return false
//}
func ExistUserByID(userId int) (isExist bool, user User) {
	if userId <= 0 {
		return false, user
	}

	db.Where("user_id = ?", userId).First(&user)
	if user.UserId > 0 {
		return true, user
	}

	return false, user
}

//func DeleteUser(id int  ) bool {
//	db. Where("id = ?", id).Delete(&Tag{})
//
//	return true
//}

func UpdateUser(id int, data interface{}) bool {
	db.Model(&User{}).Where("user_id = ?", id).Update(data)

	return true
}

/**
获取最新的用户id
*/
func GetNewUserId() int {
	var user User
	db.Select("user_id").Last(&user)
	return user.UserId + 1
}

//添加用户
func AddUser(email, password string) bool {
	//TimeFormat := "20060102 12:12:12"

	var newId = GetNewUserId()

	err := db.Create(&User{
		UserId:        newId,
		Email:         email,
		Password:      password,
		Nickname:      "昵称",
		Avatar:        "s",
		Sex:           0,
		Signature:     "...",
		RegisterTime:  time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	}).Error
	if err != nil {
		logging.Error(err)
	}
	return true
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	if scope.HasColumn("created_at") {
		scope.SetColumn("created_at", time.Now().Unix())
	}
	return nil
}
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	if scope.HasColumn("updated_at") {
		scope.SetColumn("updated_at", time.Now().Unix())
	}
	return nil
}
