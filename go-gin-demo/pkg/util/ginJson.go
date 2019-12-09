package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"net/http"
)

//注册
func GinJson(c *gin.Context, code int, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
} //注册
func GinJsonMsg(c *gin.Context, code int, msg string, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
