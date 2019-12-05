package models

import (
	"github.com/mirkowu/go-gin-demo/pkg/logging"
)

type Feedback struct {
	Model

	UserId     int    `form:"user_id" json:"user_id"`
	Title      string `form:"title"  json:"title"`
	Content    string `form:"content" json:"content"`
	Contact    string `form:"contact" json:"contact"`
	SubmitTime int64  `json:"submit_time"`
}

//添加反馈
func AddFeedback(feedback Feedback) {
	err := db.Create(&feedback).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//删除反馈
func DeleteFeedback(id int) {
	err := db.Where("id = ?", id).Delete(&Feedback{}).Error
	if err != nil {
		logging.Error(err)
	}
	return
}

//是否存在
func ExistFeedbackByID(id int) (isExist bool, feedback Feedback) {
	err := db.Where("id = ?", id).First(&feedback).Error

	if err != nil {
		logging.Error(err)
		return false, feedback
	}
	if feedback.ID > 0 {
		return true, feedback
	}
	return false, feedback

}

//更新反馈
func UpdateFeedback(id int, data interface{}) bool {
	err := db.Model(&Feedback{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

//获取指定的反馈
func GetFeedbackByID(id int) (feedback Feedback) {
	db.Where("id = ?", id).First(&feedback)
	return
}

//获取某个用户的所有列表
func GetFeedbackByUserId(userId int, page, pageSize int) (feedbacks []Feedback, total int) {
	db.Where("user_id = ?", userId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&feedbacks)
	db.Model(&Feedback{}).Count(&total)
	return
}

//获取所有列表
func GetAllFeedback(page, pageSize int) (feedbacks []Feedback, total int) {
	db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&feedbacks)
	// 获取总条数
	db.Model(&Feedback{}).Count(&total)
	return
}

//
//func (feedback *Feedback) BeforeCreate(scope *gorm.Scope) error {
//	if scope.HasColumn("created_at") {
//		scope.SetColumn("created_at", time.Now().Unix())
//	}
//	return nil
//}
//func (feedback *Feedback) BeforeUpdate(scope *gorm.Scope) error {
//	if scope.HasColumn("updated_at") {
//		scope.SetColumn("updated_at", time.Now().Unix())
//	}
//	return nil
//}
