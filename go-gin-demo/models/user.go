package models

import (
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"sync"
	"time"
)

//只做查询
type SimpleUser struct {
	UserId    int    `json:"user_id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Sex       int    `json:"sex"`
	Signature string `json:"signature"`
}

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

var lock sync.Mutex

/**
获取最新的用户id
*/
func GetNewUserId() int {
	var user User
	if err := db.Select("user_id").Last(&user).Error; err != nil {
		logging.Error(err)
	}
	return user.UserId + 1
}

//添加用户
func AddUser(email, password string) bool {
	//这里操作要保证id线程安全
	lock.Lock()
	var newId = GetNewUserId()
	err := db.Create(&User{
		UserId:        newId,
		Email:         email,
		Password:      password,
		Nickname:      "",
		Avatar:        "",
		Sex:           0,
		Signature:     "",
		RegisterTime:  time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	}).Error
	lock.Unlock()

	if err != nil {
		logging.Error(err)
	}
	return true
}

func GetUserByEmail(email string) (user User) {
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		logging.Error(err)
	}
	return
}

//func GetUserWithPwdByEmail(email string) (user User) {
//	db.Where("email = ?", email).First(&user)
//	fmt.Println(user.Password)
//	return
//}

func GetUserTotal() (count int) {
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		logging.Error(err)
	}
	return
}

func ExistUserByEmail(email string) bool {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() != "record not found" { //record not found 没查到说明不存在
			logging.Error(err)
		}
		return false
	}
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

	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		logging.Error(err)
	}
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

func UpdateUser(id int, user User) bool {
	if err := db.Model(&User{}).Where("user_id = ?", id).Update(user).Error; err != nil {
		logging.Error(err)
		return false
	}
	return true
}

//<<<<<<<<<<<< 查询别人的信息 视情况给字段
var sqlOtherUser = "user_id,nickname,avatar,sex,signature"

func GetUserByID(userId int) (isExist bool, user SimpleUser) {
	if userId <= 0 {
		return false, user
	}
	if err := db.Table("user").Select(sqlOtherUser).Where("user_id = ?", userId).First(&user).Error; err != nil {
		logging.Error(err)
	}
	if user.UserId > 0 {
		return true, user
	}

	return false, user
}

func GetAllUser(pageSize, offset int) (user []SimpleUser) {
	if err := db.Table("user").Select(sqlOtherUser).Limit(pageSize).Offset(offset).Find(&user).Error; err != nil {
		logging.Error(err)
	}
	return
}

//查询别人的信息 >>>>>>>>>>>

//
//func (user *User) BeforeCreate(scope *gorm.Scope) error {
//	if scope.HasColumn("created_at") {
//		scope.SetColumn("created_at", time.Now().Unix())
//	}
//	return nil
//}
//func (user *User) BeforeUpdate(scope *gorm.Scope) error {
//	if scope.HasColumn("updated_at") {
//		scope.SetColumn("updated_at", time.Now().Unix())
//	}
//	return nil
//}
