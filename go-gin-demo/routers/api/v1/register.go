package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

const (
	TypeCaptchaRegister      = 1
	TypeCaptchaResetPassword = 2
)

//获取验证码
func GetCaptcha(c *gin.Context) {
	email := c.Query("email")
	captchaType := com.StrTo(c.DefaultQuery("type", "0")).MustInt()

	var data string
	code := e.ERROR_EMAIL
	if util.CheckEmail(email) {

		if captchaType == TypeCaptchaRegister {
			if models.ExistUserByEmail(email) {
				code = e.ERROR_EXIST_EMAIL
			} else {
				code = e.SUCCESS
				data = "123456"
			}
		} else if captchaType == TypeCaptchaResetPassword {
			if !models.ExistUserByEmail(email) {
				code = e.ERROR_NOT_EXIST_EMAIL
			} else {
				code = e.SUCCESS
				data = "123456"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//注册
func Register(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	var data interface{}
	code := e.ERROR_EMAIL_PASSWORD
	if util.CheckEmailAndPwd(email, password) {
		if models.ExistUserByEmail(email) {
			code = e.ERROR_EXIST_EMAIL
		} else {
			models.AddUser(email, password)
			code = e.SUCCESS
			user := models.GetUserByEmail(email)
			data = user
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//登录
func Login(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	var data interface{}
	code := e.ERROR_EMAIL_PASSWORD
	if util.CheckEmailAndPwd(email, password) {
		if models.ExistUserByEmail(email) {
			user := models.GetUserByEmail(email)
			if password == user.Password {
				data = user
			} else {
				code = e.ERROR_PASSWORD
			}
		} else {
			code = e.ERROR_NOT_EXIST_EMAIL
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
