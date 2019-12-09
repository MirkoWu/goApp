package util

import (
	"github.com/astaxie/beego/validation"
	"github.com/mirkowu/go-gin-demo/models"
	"regexp"
)

const (
	REG_PASSWORD string = "^[a-z0-9_-]{6,18}$"
)

//检测邮箱密码是否非法
func CheckEmail(email string) string {
	valid := validation.Validation{}
	valid.Required(email, "email").Message("邮箱不能为空")
	valid.Email(email, "email").Message("邮箱不合法")

	if valid.HasErrors() {
		return valid.Errors[0].Message
	}
	return ""
	//return !valid.HasErrors() ,   valid.Errors[0].Message
}

//检测邮箱密码是否非法
func CheckPwd(password string) string {
	valid := validation.Validation{}
	valid.Required(password, "password").Message("密码不能为空")
	valid.Match(password, regexp.MustCompile(REG_PASSWORD), "password").Message("密码不合法")

	if valid.HasErrors() {
		return valid.Errors[0].Message
	}
	return ""
}

//检测更新密码
func CheckUpdatePwd(oldPassword, newPassword string) string {
	valid := validation.Validation{}
	valid.Required(oldPassword, "password").Message("旧密码不能为空")
	valid.Match(oldPassword, regexp.MustCompile(REG_PASSWORD), "password").Message("旧密码不合法")
	valid.Required(newPassword, "password").Message("新密码不能为空")
	valid.Match(newPassword, regexp.MustCompile(REG_PASSWORD), "password").Message("新密码不合法")
	if valid.HasErrors() {
		return valid.Errors[0].Message
	}
	return ""
}

//检测邮箱密码是否非法
func CheckEmailAndPwd(email, password string) string {

	valid := validation.Validation{}
	valid.Required(email, "email").Message("邮箱不能为空")
	valid.Email(email, "email").Message("邮箱不合法")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Match(password, regexp.MustCompile(REG_PASSWORD), "password").Message("密码不合法")

	if valid.HasErrors() {
		return valid.Errors[0].Message
	}
	return ""
}

//检测APP提交
func CheckAppInfo(data models.AppShow) string {

	valid := validation.Validation{}
	valid.Required(data.Name, "Name").Message("名称不能为空")
	valid.Required(data.Intro, "Intro").Message("介绍不能为空")
	valid.Required(data.AppPn, "AppPn").Message("包名不能为空")
	valid.Required(data.AppVersion, "AppVersion").Message("版本不能为空")
	valid.Min(data.AppVersionCode, 1, "AppVersion").Message("版本号必须大于0")

	if valid.HasErrors() {
		return valid.Errors[0].Message
	}
	return ""
}
