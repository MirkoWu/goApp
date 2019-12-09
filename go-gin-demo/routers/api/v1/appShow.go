package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
)

//添加
func AddAppShow(c *gin.Context) {
	//userId := util.GetUserId(c)

	var data models.AppShow
	code := e.ERROR_NOT_EXIST_USER
	if err := c.ShouldBind(&data); err != nil {
		code = e.ERROR
	} else {
		if msg := util.CheckAppInfo(data); msg == "" {
			//	if isExist, _ := models.ExistUserByID(userId); isExist {
			models.AddAppShow(data)
			code = e.SUCCESS
			//	}
		} else {
			code = e.INVALID_PARAMS
			util.GinJsonMsg(c, code, msg, nil)
			return
		}
	}

	util.GinJson(c, code, nil)
}

////查询所有
//func GetAllAppTab(c *gin.Context) {
//	//userId := util.GetUserId(c)
//	//pageSize, offset := util.GetPageByPost(c)
//
//	var list []models.AppTab
//	code := e.ERROR_NOT_EXIST_USER
//	//if isExist, _ := models.ExistUserByID(userId); isExist {
//		list, _ = models.GetAllShowAppTab()
//		code = e.SUCCESS
//	//}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": list,
//	})
//}
//查询所有显示的列表
func GetAllShowAppList(c *gin.Context) {
	//userId := util.GetUserId(c)
	pageSize, offset := util.GetPageByPost(c)

	var list []models.AppShow
	code := e.ERROR_NOT_EXIST_USER
	//if isExist, _ := models.ExistUserByID(userId); isExist {
	list, _ = models.GetAllShowAppList(pageSize, offset)
	code = e.SUCCESS
	//}

	util.GinJson(c, code, list)
}

//
////更新
//func UpdateFeedback(c *gin.Context) {
//	userId := util.GetUserId(c)
//
//	id := com.StrTo(c.PostForm("id")).MustInt()
//	title := c.PostForm("title")
//	content := c.PostForm("content")
//	contact := c.PostForm("contact")
//
//	var list []models.Feedback
//	code := e.ERROR_NOT_EXIST_USER
//	if isExist, _ := models.ExistUserByID(userId); isExist {
//		if isExist, feedback := models.ExistFeedbackByID(id); isExist {
//			if title != "" {
//				feedback.Title = title
//			}
//			if content != "" {
//				feedback.Content = content
//			}
//			if contact != "" {
//				feedback.Contact = contact
//			}
//
//			models.UpdateFeedback(id, feedback)
//			code = e.SUCCESS
//		} else {
//			code = e.ERROR_NOT_EXIST
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": list,
//	})
//}
//
////删除反馈
//func DeleteFeedback(c *gin.Context) {
//	userId := util.GetUserId(c)
//	id := com.StrTo(c.PostForm("id")).MustInt()
//
//	code := e.ERROR_NOT_EXIST_USER
//	if isExist, _ := models.ExistUserByID(userId); isExist {
//		if isExist, _ := models.ExistFeedbackByID(id); isExist {
//			models.DeleteFeedback(id)
//			code = e.SUCCESS
//		} else {
//			code = e.ERROR_NOT_EXIST
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": nil,
//	})
//}
