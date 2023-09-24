package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCore 创建 jwt 结构体
type JwtCore[T any] struct {
	VerifyKey []byte // 加密密钥
}

// GenerateTokenParameters 生成token 必备参数
type GenerateTokenParameters[T any] struct {
	Data T `json:"data"`
	jwt.RegisteredClaims
}

// Create 生成并且返回 token
func (receiver JwtCore[T]) Create(opt GenerateTokenParameters[T]) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, opt).SignedString(receiver.VerifyKey)
}

// Parse 解析 token 返回一个 GenerateTokenParameters 结构体
func (receiver JwtCore[T]) Parse(token string) (*GenerateTokenParameters[T], error) {
	parse, err := jwt.ParseWithClaims(token, &GenerateTokenParameters[T]{}, func(token *jwt.Token) (interface{}, error) {
		return receiver.VerifyKey, nil
	})

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, nil
	}

	if claims, ok := parse.Claims.(*GenerateTokenParameters[T]); ok && parse.Valid {
		return claims, nil
	}

	return nil, nil
}
