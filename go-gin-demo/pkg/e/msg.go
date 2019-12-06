package e

var MsgFlags = map[int]string{
	SUCCESS:                        "success",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_EMAIL_PASSWORD:            "邮箱或密码错误",
	ERROR_EXIST_EMAIL:               "邮箱已存在",
	ERROR_NOT_EXIST_EMAIL:           "邮箱不存在",
	ERROR_PASSWORD:                  "密码错误",
	ERROR_OLD_PASSWORD:              "旧密码错误",
	ERROR_NOT_SAME_OLD_NEW_PASSWORD: "旧密码和新密码不一致",
	ERROR_NOT_EXIST_USER:            "用户不存在",
	ERROR_NOT_EXIST_USER_BY_QUREY:   "查询的用户不存在",

	ERROR_UPLOAD_FILE:   "上传失败",
	ERROR_FILE_OUT_SIZE: "图片过大",

	ERROR_NOT_EMPTY:         "内容不能为空",
	ERROR_NOT_EMPTY_TITLE:   "标题不能为空",
	ERROR_NOT_EMPTY_CONTENT: "内容不能为空",
	ERROR_NOT_EXIST:         "内容不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
