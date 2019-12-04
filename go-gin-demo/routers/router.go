package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/middleware/jwt"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
	"github.com/mirkowu/go-gin-demo/routers/api"
	v1 "github.com/mirkowu/go-gin-demo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	//验证
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//上传多张图片
		apiv1.POST("/upload_files", v1.UploadFiles)

		//注册
		apiv1.POST("/register", v1.Register)
		//登录
		apiv1.POST("/login", v1.Login)
		//获取验证码
		apiv1.POST("/get_captcha", v1.GetCaptcha)

		//更新密码
		apiv1.POST("/update_password", v1.UpdatePassword)
		//更新用户信息
		apiv1.POST("/update_user_info", v1.UpdateUserInfo)
		//更新头像
		apiv1.POST("/update_avatar", v1.UpdateAvatar)
	}

	return r
}
