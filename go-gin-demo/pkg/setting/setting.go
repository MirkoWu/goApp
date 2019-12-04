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
	DB_USER            = "root"
	DB_PASSWORD        = "root"
	//127.0.0.1:3306
	DB_HOST         = "127.0.0.1:3306"
	DB_NAME         = "user"
	DB_TABLE_PREFIX = ""

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
