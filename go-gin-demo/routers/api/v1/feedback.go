package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/unknwon/com"
	"net/http"
)

//添加
func AddFeedback(c *gin.Context) {
	//userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt()
	userId := c.GetInt("user_id") //token中取

	var feedback models.Feedback
	code := e.ERROR_NOT_EXIST_USER
	if err := c.ShouldBind(&feedback); err != nil {
		code = e.ERROR
	} else {
		if feedback.Title == "" || feedback.Content == "" {
			code = e.ERROR_NOT_EMPTY
		} else {
			if isExist, _ := models.ExistUserByID(userId); isExist {
				feedback.UserId = userId
				models.AddFeedback(feedback)
				code = e.SUCCESS
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

//获取
func GetAllFeedback(c *gin.Context) {
	//userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt()
	userId := c.GetInt("user_id") //token中取

	page := com.StrTo(c.PostForm("page")).MustInt()
	pageSize := com.StrTo(c.PostForm("pageSize")).MustInt()

	var list []models.Feedback

	code := e.ERROR_NOT_EXIST_USER
	if isExist, _ := models.ExistUserByID(userId); isExist {
		list, _ = models.GetAllFeedback(page, pageSize)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": list,
	})
}

//获取
func UpdateFeedback(c *gin.Context) {
	//userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt()
	userId := c.GetInt("user_id") //token中取

	id := com.StrTo(c.PostForm("id")).MustInt()
	title := c.PostForm("title")
	content := c.PostForm("content")
	contact := c.PostForm("contact")

	var list []models.Feedback

	code := e.ERROR_NOT_EXIST_USER
	if isExist, _ := models.ExistUserByID(userId); isExist {
		if isExist, feedback := models.ExistFeedbackByID(id); isExist {
			if title != "" {
				feedback.Title = title
			}
			if content != "" {
				feedback.Content = content
			}
			if contact != "" {
				feedback.Contact = contact
			}

			models.UpdateFeedback(id, feedback)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": list,
	})
}
