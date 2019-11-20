package setting

import (
	"time"
)

var (
	RunMode string = "debug"

	HTTPPort     int           = 8080
	ReadTimeout  time.Duration = 60 * time.Second
	WriteTimeout time.Duration = 60 * time.Second

	PageSize  int    = 10
	JwtSecret string = "23347$040412"

	DB_TYPE     string = "mysql"
	DB_USER            = "数据库账号"
	DB_PASSWORD        = "数据库密码"
	//127.0.0.1:3306
	DB_HOST         = "数据库IP:数据库端口号"
	DB_NAME         = "blog"
	DB_TABLE_PREFIX = "blog_"

	//logging
	RuntimeRootPath = "runtime/"
	LogSavePath     = "logs/"
	LogSaveName     = "log"
	LogFileExt      = "log"
	TimeFormat      = 20060102
	//image
	ImagePrefixUrl = "http://127.0.0.1:8000"
	ImageSavePath  = "upload/images/"
	ImageMaxSize   = 5
	ImageAllowExts = " .jpg,.jpeg,.png"
)
