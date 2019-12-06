package util

import (
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) int {
	//userId := com.StrTo(c.DefaultQuery("user_id", "0")).MustInt()
	userId := c.GetInt("user_id") //token中取

	return userId
}
