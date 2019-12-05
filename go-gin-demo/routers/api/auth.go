package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/models"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt()
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if userId > 0 {
		isExist := models.CheckAuth(userId)
		if isExist {
			token, err := util.GenerateToken(userId)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
