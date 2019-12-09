package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"github.com/unknwon/com"
	"time"
)

const (
	TypeCaptchaRegister      = 1
	TypeCaptchaResetPassword = 2
)

//获取验证码
func GetCaptcha(c *gin.Context) {
	email := c.PostForm("email")
	captchaType := com.StrTo(c.DefaultQuery("type", "0")).MustInt()

	var data string
	code := e.ERROR_EMAIL
	if msg := util.CheckEmail(email); msg == "" {
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
		} else {
			code = e.INVALID_PARAMS
		}

		util.GinJson(c, code, data)
	} else {
		util.GinJsonMsg(c, code, msg, data)
	}

}

//注册
func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var data interface{}
	code := e.ERROR_EMAIL_PASSWORD
	if msg := util.CheckEmailAndPwd(email, password); msg == "" {
		if models.ExistUserByEmail(email) {
			code = e.ERROR_EXIST_EMAIL
		} else {
			models.AddUser(email, password)
			user := models.GetUserByEmail(email)

			user.Token, _ = util.GenerateToken(user.UserId) //token
			models.UpdateUser(user.UserId, user)            //更新

			data = user
			code = e.SUCCESS
		}
		util.GinJson(c, code, data)
	} else {
		util.GinJsonMsg(c, code, msg, data)
	}

}

//登录
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var data interface{}
	code := e.ERROR_EMAIL_PASSWORD
	if msg := util.CheckEmailAndPwd(email, password); msg == "" {
		if models.ExistUserByEmail(email) {
			user := models.GetUserByEmail(email)
			if password == user.Password {
				user.LastLoginTime = time.Now().Unix()          //登录时间
				user.Token, _ = util.GenerateToken(user.UserId) //token
				models.UpdateUser(user.UserId, user)            //更新

				data = user
				code = e.SUCCESS
			} else {
				code = e.ERROR_PASSWORD
			}
		} else {
			code = e.ERROR_NOT_EXIST_EMAIL
		}
		util.GinJson(c, code, data)
	} else {
		util.GinJsonMsg(c, code, msg, data)
	}

}
