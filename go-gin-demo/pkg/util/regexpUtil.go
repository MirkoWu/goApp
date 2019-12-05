package util

import (
	"github.com/astaxie/beego/validation"
	"regexp"
)

const (
	REG_PASSWORD string = "^[a-z0-9_-]{6,18}$"
)

//检测邮箱密码是否非法
func CheckEmail(email string) bool {
	valid := validation.Validation{}
	valid.Required(email, "email").Message("邮箱不能为空")
	valid.Email(email, "email").Message("邮箱不合法")
	return !valid.HasErrors()
}

//检测邮箱密码是否非法
func CheckPwd(password string) bool {
	valid := validation.Validation{}
	valid.Required(password, "password").Message("密码不能为空")
	valid.Match(password, regexp.MustCompile(REG_PASSWORD), "password").Message("密码不合法")
	return !valid.HasErrors()
}

//检测邮箱密码是否非法
func CheckEmailAndPwd(email, password string) bool {
	return CheckEmail(email) && CheckPwd(password)
	//valid := validation.Validation{}
	//valid.Required(email, "email").Message("邮箱不能为空")
	//valid.Email(email, "email").Message("邮箱不合法")
	//valid.Required(password, "password").Message("密码不能为空")
	//valid.Match(password, Reg_password, "password").Message("密码不合法")
	//
	//return !valid.HasErrors()
}
