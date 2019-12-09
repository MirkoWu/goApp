package models

import (
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"sync"
)

//app 推荐
type AppShow struct {
	Model

	AppId          int    `form:"app_id" json:"app_id"`
	Name           string `form:"name" json:"name"`
	Intro          string `form:"intro" json:"intro"`
	Screenshots    string `form:"screenshots" json:"screenshots"`
	LinkUrl        string `form:"link_url" json:"link_url"`
	ApkUrl         string `form:"apk_url" json:"apk_url"`
	ApkSize        string `form:"apk_size" json:"apk_size"`
	AppVersion     string `form:"app_version" json:"app_version"`
	AppVersionCode int    `form:"app_version_code" json:"app_version_code"`
	AppPn          string `form:"app_pn" json:"app_pn"` //包名
	IsShow         int    `form:"is_show" json:"-"`     //是否显示
}

var lock4 sync.Mutex

//获取的id
func GetNewAppId() int {
	var data AppShow
	if err := db.Select("app_id").Last(&data).Error; err != nil {
		logging.Error(err)
	}
	return data.AppId + 1
}

//添加
func AddAppShow(data AppShow) {
	lock4.Lock()
	data.AppId = GetNewAppId() //更新id
	err := db.Create(&data).Error
	lock4.Unlock()

	if err != nil {
		logging.Error(err)
	}

	return
}

//删除
func DeleteApp(id int) {
	err := db.Where("app_id = ?", id).Delete(&AppShow{}).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//是否存在
func ExistAppByID(id int) (isExist bool, data AppShow) {
	err := db.Where("app_id = ?", id).First(&data).Error

	if err != nil {
		logging.Error(err)
		return false, data
	}
	if data.ID > 0 {
		return true, data
	}
	return false, data

}

//更新
func UpdateAppShow(id int, data AppShow) bool {
	err := db.Model(&AppShow{}).Where("app_id = ?", id).Update(data).Error
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

//获取所有列表
func GetAllApp(pageSize, offset int) (list []AppShow, total int) {
	db.Limit(pageSize).Offset(offset).Find(&list)
	// 获取总条数
	db.Model(&AppShow{}).Count(&total)
	return
}

//获取所有列表
func GetAllShowAppList(pageSize, offset int) (list []AppShow, total int) {
	db.Where("is_show = ?", 1).Limit(pageSize).Offset(offset).Find(&list)
	// 获取总条数
	db.Model(&AppShow{}).Count(&total)
	return
}
