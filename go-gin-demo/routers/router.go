package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/middleware/jwt"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
	v1 "github.com/mirkowu/go-gin-demo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//验证
	//r.GET("/auth", api.GetAuth)

	loginApi := r.Group("/api")
	loginApi.Use()
	{
		//这些不需要token
		//注册
		loginApi.POST("/register", v1.Register)
		//获取验证码
		loginApi.POST("/get_captcha", v1.GetCaptcha)
		//登录
		loginApi.POST("/login", v1.Login)

		//获取用户信息
		loginApi.POST("/get_user_info", v1.GetUserInfo)
		//更新头像
		loginApi.POST("/update_avatar", v1.UpdateAvatar)

		//上传多张图片
		loginApi.POST("/upload_files", v1.UploadFiles)

	}

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
		//获取用户信息
		apiv1.POST("/get_user_info", v1.GetUserInfo)
		//更新密码
		apiv1.POST("/update_password", v1.UpdatePassword)
		//更新用户信息
		apiv1.POST("/update_user_info", v1.UpdateUserInfo)
		//更新头像
		apiv1.POST("/update_avatar", v1.UpdateAvatar)

		//提交反馈
		apiv1.POST("/submit_feedback", v1.AddFeedback)
		//获取反馈列表
		apiv1.POST("/get_feedback_list", v1.GetAllFeedback)
		//更新反馈
		apiv1.POST("/update_feedback", v1.UpdateFeedback)

	}

	return r
}
