package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
func GetPageByGet(c *gin.Context) (pageSize, offset int) {

	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()
	pageSize = com.StrTo(c.DefaultQuery("pageSize", "0")).MustInt()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = setting.PageSize
	}
	offset = (page - 1) * setting.PageSize

	return
}

func GetPageByPost(c *gin.Context) (pageSize, offset int) {
	page := com.StrTo(c.PostForm("page")).MustInt()
	pageSize = com.StrTo(c.PostForm("pageSize")).MustInt()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = setting.PageSize
	}
	offset = (page - 1) * setting.PageSize
	return
}
