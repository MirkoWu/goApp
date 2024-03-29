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

	noTokenApi := r.Group("/api/v1")
	noTokenApi.Use()
	{
		//这些不需要token
		//注册
		noTokenApi.POST("/register", v1.Register)
		//获取验证码
		noTokenApi.POST("/get_captcha", v1.GetCaptcha)
		//登录
		noTokenApi.POST("/login", v1.Login)

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

		//获取用户列表
		apiv1.POST("/get_user_list", v1.GetAllUser)
		//获取用户信息
		apiv1.POST("/get_user_info", v1.GetUserInfo)
		//获取指定用户信息
		apiv1.POST("/get_user_info_by_id", v1.GetUserInfoByID)
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
		//删除反馈
		apiv1.POST("/delete_feedback", v1.DeleteFeedback)

		//添加tab
		apiv1.POST("/add_app_tab", v1.AddAppTab)
		//tab列表
		apiv1.POST("/get_app_tab_list", v1.GetAllShowAppTab)

		//添加app
		apiv1.POST("/add_app_show", v1.AddAppShow)
		//app列表
		apiv1.POST("/get_app_show_list", v1.GetAllShowAppList)

	}

	//admin

	return r
}
