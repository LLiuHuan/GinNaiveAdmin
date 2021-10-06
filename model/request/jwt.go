package request

import (
	"GinNaiveAdmin/model"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	AccessToken string
	UserInfo    model.JwtUser
	//BufferTime   int64   // 缓冲时间 暂时去掉
	jwt.StandardClaims
}
