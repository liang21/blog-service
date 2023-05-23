package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lgl/blog-service/global"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	// 生成token
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	// 创建claim
	claims := Claims{
		AppKey:    appKey,
		AppSecret: appSecret,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 发布时间
			IssuedAt: nowTime.Unix(),
			// 发布者
			Issuer: "blog-service",
			// 主题
			Subject: "user token",
		},
	}
	// 创建token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err

}

func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验token的签名
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		// 校验token的有效性
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	// 校验token的有效性
	return nil, err
}
