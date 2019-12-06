package e

const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_EMAIL                     = 10010
	ERROR_EMAIL_PASSWORD            = 10011
	ERROR_EXIST_EMAIL               = 10012
	ERROR_NOT_EXIST_EMAIL           = 10013
	ERROR_PASSWORD                  = 10014
	ERROR_OLD_PASSWORD              = 10015
	ERROR_NOT_SAME_OLD_NEW_PASSWORD = 10016
	ERROR_NOT_EXIST_USER            = 10017
	ERROR_NOT_EXIST_USER_BY_QUREY   = 10018

	ERROR_UPLOAD_FILE   = 10020
	ERROR_FILE_OUT_SIZE = 10021

	ERROR_NOT_EMPTY         = 10030
	ERROR_NOT_EMPTY_TITLE   = 10031
	ERROR_NOT_EMPTY_CONTENT = 10032
	ERROR_NOT_EXIST         = 10033
)
