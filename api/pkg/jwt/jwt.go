package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("hello musk")

type Claims struct {
	Userid int
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userid int) (tokenString string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	claims := Claims{
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	tokenString, err = token.SignedString(secret)
	return
}

func GenerateRefreshToken(userid int) (tokenString string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	claims := Claims{
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	tokenString, err = token.SignedString(secret)
	return
}

// ParseToken 解析token
// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
