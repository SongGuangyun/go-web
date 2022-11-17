package utils

import (
	"github.com/Songguangyun/go-web/internal/constant"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	SigningKey []byte
}

// NewJWT 构造方法
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ConfigServer.Jwt.SigningKey),
	}
}

type CustomClaims struct {
	UID string
	jwt.RegisteredClaims
}

func CreateClaims(uid string, ttl int64) CustomClaims {
	now := time.Now()
	before := now.Add(-time.Minute * 5)
	//fmt.Println(time.Now().Unix())
	return CustomClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                                     // 发布时间
			NotBefore: jwt.NewNumericDate(before),                                  // JWT 之前的时间 不得接受处理
		},
	}
}

func (j *JWT) CreateToken(userID string) (string, error) {
	claims := CreateClaims(userID, global.ConfigServer.Jwt.ExpireTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 返回值类似：0xc000447de0
	tokenString, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) GetClaimFromToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, keyFunc(j.SigningKey))
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, constant.TokenMalformedMsg
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, constant.TokenExpiredMsg
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, constant.TokenNotValidYetMsg
			} else {
				return nil, constant.TokenUnknownMsg
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, constant.TokenUnknownMsg

	} else {
		return nil, constant.TokenUnknownMsg
	}
}

func keyFunc(signingKey []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}
}
