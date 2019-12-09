package models

import (
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"sync"
)

type AppTab struct {
	Model

	TabId  int    `form:"tab_id" json:"tab_id"`
	Title  string `form:"title" json:"title"`
	Type   int    `form:"type"  json:"type"`
	IsShow int    `form:"is_show" json:"-"`
}

var lock3 sync.Mutex

//获取的id
func GetNewTabId() int {
	var data AppTab
	if err := db.Select("tab_id").Last(&data).Error; err != nil {
		logging.Error(err)
	}
	return data.TabId + 1
}

//添加
func AddAppTab(data AppTab) {
	lock3.Lock()
	data.TabId = GetNewTabId() //更新id
	err := db.Create(&data).Error
	lock3.Unlock()

	if err != nil {
		logging.Error(err)
	}
	return
}

//删除
func DeleteAppTab(id int) {
	err := db.Where("tab_id = ?", id).Delete(&AppTab{}).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//是否存在
func ExistAppTabByID(id int) (isExist bool, data AppTab) {
	err := db.Where("tab_id = ?", id).First(&data).Error

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
func UpdateAppTab(id int, data AppTab) bool {
	err := db.Model(&AppTab{}).Where("tab_id = ?", id).Update(data).Error
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

//获取所有列表
func GetAllAppTab(pageSize, offset int) (list []AppTab, total int) {
	db.Limit(pageSize).Offset(offset).Find(&list)
	// 获取总条数
	db.Model(&AppTab{}).Count(&total)
	return
}

//获取所有列表
func GetAllShowAppTab() (list []AppTab, total int) {
	db.Where("is_show = ?", 1).Find(&list)
	// 获取总条数
	db.Model(&AppTab{}).Count(&total)
	return
}
