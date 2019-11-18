module github.com/mirkowu/go-gin-demo

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20191118013547-6254a7c3cac6 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace (
	github.com/mirkowu/go-gin-demo/conf => ../GoCode/application/go-gin-demo/conf
	github.com/mirkowu/go-gin-demo/middleware => ../GoCode/application/go-gin-demo/middleware
	github.com/mirkowu/go-gin-demo/middleware/jwt => ../GoCode/application/go-gin-demo/middleware/jwt
	github.com/mirkowu/go-gin-demo/models => ../GoCode/application/go-gin-demo/models
	github.com/mirkowu/go-gin-demo/pkg/e => ../GoCode/application/go-gin-demo/pkg/e
	github.com/mirkowu/go-gin-demo/pkg/logging => ../GoCode/application/go-gin-demo/pkg/logging
	github.com/mirkowu/go-gin-demo/pkg/setting => ../GoCode/application/go-gin-demo/pkg/setting
	github.com/mirkowu/go-gin-demo/pkg/util => ../GoCode/application/go-gin-demo/pkg/util
	github.com/mirkowu/go-gin-demo/routers => ../GoCode/application/go-gin-demo/routers
	//github.com/mirkowu/go-gin-demo/routers/api/v1 => ../GoCode/application/go-gin-demo/routers/api/v1
	github.com/mirkowu/go-gin-demo/runtime => ../GoCode/application/go-gin-demo/runtime

)
