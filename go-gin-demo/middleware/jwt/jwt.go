package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("token")
		//token := c.Query("token")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				if claims.UserId > 0 {
					//解析 user_id  接口直接用，客户端只需传token就好了
					c.Set("user_id", claims.UserId)
				} else {
					code = e.ERROR_NOT_EXIST_USER
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理
		c.Next()
		//并将解析出的信息传递下去
		//c.Set("user_id",1)
	}
}
