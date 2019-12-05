package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	UserId int `json:"user_id"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

/**
生成token
*/
func GenerateToken(userId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour) //有效期7天

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-demo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		logging.Error(err)
	}
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
