package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"github.com/unknwon/com"
	"net/http"
	"regexp"
)

//修改密码
func UpdatePassword(c *gin.Context) {
	userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt64()
	oldPassword := c.Query("old_password")
	newPassword := c.Query("new_password")

	valid := validation.Validation{}
	valid.Required(oldPassword, "password").Message("旧密码不能为空")
	valid.Match(oldPassword, regexp.MustCompile(util.REG_PASSWORD), "password").Message("旧密码不合法")
	valid.Required(newPassword, "password").Message("新密码不能为空")
	valid.Match(newPassword, regexp.MustCompile(util.REG_PASSWORD), "password").Message("新密码不合法")

	code := e.ERROR_EMAIL_PASSWORD
	if !valid.HasErrors() {
		if oldPassword == newPassword {
			if models.ExistUserByID(userId) {
				user := models.GetUserByID(userId)
				if oldPassword == user.Password {
					user.Password = newPassword
					models.EditUser(userId, user)
					//修改成功 之后一般还会重置下token
					code = e.SUCCESS
				} else {
					code = e.ERROR_OLD_PASSWORD
				}
			} else {
				code = e.ERROR_NOT_EXIST_USER
			}
		} else {
			code = e.ERROR_NOT_SAME_OLD_NEW_PASSWORD
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

//修改密码
func UpdateUserInfo(c *gin.Context) {
	userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt64()
	nickname := c.Query("nickname")
	ageStr := c.Query("age")
	signature := c.Query("signature")

	code := e.SUCCESS
	if models.ExistUserByID(userId) {
		user := models.GetUserByID(userId)
		if nickname != "" {
			user.Nickname = nickname
		}
		if signature != "" {
			user.Signature = signature
		}
		if ageStr != "" {
			user.Age = com.StrTo(ageStr).MustInt()
		}
		models.EditUser(userId, user)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}

//更新头像
func UpdateAvatar(c *gin.Context) {
	userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt64()

	code := e.SUCCESS
	if models.ExistUserByID(userId) {
		user := models.GetUserByID(userId)
		code2, data, _ := UploadFile(c)
		code = code2
		if code == e.SUCCESS {
			user.Avatar = data[0]
			models.EditUser(userId, user)
		}
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": nil,
	})
}
