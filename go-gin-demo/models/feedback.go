package models

import (
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"time"
)

type Feedback struct {
	Model

	FeedbackId int    `form:"feedback_id" json:"feedback_id"`
	UserId     int    `form:"user_id" json:"user_id"`
	Title      string `form:"title"  json:"title"`
	Content    string `form:"content" json:"content"`
	Contact    string `form:"contact" json:"contact"`
	SubmitTime int64  `json:"submit_time"`
}

//获取的id
func GetNewFeedbackId() int {
	var data Feedback
	if err := db.Select("feedback_id").Last(&data).Error; err != nil {
		logging.Error(err)
	}
	return data.FeedbackId + 1
}

//添加反馈
func AddFeedback(data Feedback) {
	data.FeedbackId = GetNewFeedbackId() //更新id
	data.SubmitTime = time.Now().Unix()
	err := db.Create(&data).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//删除反馈
func DeleteFeedback(id int) {
	err := db.Where("feedback_id = ?", id).Delete(&Feedback{}).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//是否存在
func ExistFeedbackByID(id int) (isExist bool, data Feedback) {
	err := db.Where("feedback_id = ?", id).First(&data).Error

	if err != nil {
		logging.Error(err)
		return false, data
	}
	if data.ID > 0 {
		return true, data
	}
	return false, data

}

//更新反馈
func UpdateFeedback(id int, data Feedback) bool {
	err := db.Model(&Feedback{}).Where("feedback_id = ?", id).Update(data).Error
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

//获取指定的反馈
func GetFeedbackByID(id int) (data Feedback) {
	db.Where("feedback_id = ?", id).First(&data)
	return
}

//获取某个用户的所有列表
func GetFeedbackByUserId(userId int, page, pageSize int) (list []Feedback, total int) {
	db.Where("user_id = ?", userId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&list)
	db.Model(&Feedback{}).Count(&total)
	return
}

//获取所有列表
func GetAllFeedback(pageSize, offset int) (list []Feedback, total int) {
	db.Limit(pageSize).Offset(offset).Find(&list)
	// 获取总条数
	db.Model(&Feedback{}).Count(&total)
	return
}

//
//func (feedback *Feedback) BeforeCreate(scope *gorm.Scope) error {
//	//if scope.HasColumn("created_at") {
//	//	scope.SetColumn("created_at", time.Now().Unix())
//	//}
//	if scope.HasColumn("deleted_at") {
//		scope.SetColumn("deleted_at", nil)
//	}
//	return nil
//}

//func (feedback *Feedback) BeforeUpdate(scope *gorm.Scope) error {
//	if scope.HasColumn("updated_at") {
//		scope.SetColumn("updated_at", time.Now().Unix())
//	}
//	return nil
//}
