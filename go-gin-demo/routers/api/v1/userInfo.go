package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"github.com/unknwon/com"
)

//登录
func GetUserInfo(c *gin.Context) {
	userId := util.GetUserId(c)

	var data interface{}
	code := e.ERROR_NOT_EXIST_USER
	if isExist, user := models.ExistUserByID(userId); isExist {
		data = user
		code = e.SUCCESS
	}

	util.GinJson(c, code, data)
}

//查询指定用户
func GetUserInfoByID(c *gin.Context) {
	userId := util.GetUserId(c)
	targetId := com.StrTo(c.PostForm("target_id")).MustInt()

	var data interface{}
	code := e.ERROR_NOT_EXIST_USER
	if isExist, _ := models.ExistUserByID(userId); isExist {
		//查询别人的信息
		if isExist, user := models.GetUserByID(targetId); isExist {
			data = user
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_USER_BY_QUREY
		}
	}

	util.GinJson(c, code, data)
}

//获取所有用户
func GetAllUser(c *gin.Context) {
	userId := util.GetUserId(c)
	pageSize, offset := util.GetPageByPost(c)

	var data interface{}
	code := e.ERROR_NOT_EXIST_USER
	if isExist, _ := models.ExistUserByID(userId); isExist {
		data = models.GetAllUser(pageSize, offset)
		code = e.SUCCESS
	}

	util.GinJson(c, code, data)
}

//修改密码
func UpdatePassword(c *gin.Context) {
	userId := c.GetInt("user_id") //token中取
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	code := e.ERROR_EMAIL_PASSWORD
	if msg := util.CheckUpdatePwd(oldPassword, newPassword); msg == "" {
		if oldPassword == newPassword {
			if isExist, user := models.ExistUserByID(userId); isExist {

				if oldPassword == user.Password {
					user.Password = newPassword
					models.UpdateUser(userId, user)
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
		util.GinJson(c, code, nil)
	} else {
		util.GinJsonMsg(c, code, msg, nil)
	}
}

//修改密码
func UpdateUserInfo(c *gin.Context) {
	userId := c.GetInt("user_id") //token中取
	nickname := c.PostForm("nickname")
	sexStr := c.PostForm("sex")
	signature := c.PostForm("signature")

	code := e.SUCCESS
	if isExist, user := models.ExistUserByID(userId); isExist {

		if nickname != "" {
			user.Nickname = nickname
		}
		if signature != "" {
			user.Signature = signature
		}
		if sexStr != "" {
			user.Sex = com.StrTo(sexStr).MustInt()
		}
		models.UpdateUser(userId, user)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}

	util.GinJson(c, code, nil)

}

//更新头像
func UpdateAvatar(c *gin.Context) {
	userId := c.GetInt("user_id") //token中取

	code := e.SUCCESS
	if isExist, user := models.ExistUserByID(userId); isExist {
		code2, data, _ := UploadFile(c)
		code = code2
		if code == e.SUCCESS {
			user.Avatar = data[0]
			models.UpdateUser(userId, user)
		}
	} else {
		code = e.ERROR_NOT_EXIST_USER
	}

	util.GinJson(c, code, nil)
}
