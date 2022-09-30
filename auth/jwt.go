package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JwtAuth struct {
	sign   []byte        //默认16位
	expire time.Duration //默认 1小时
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{
		expire: time.Hour,
		sign:   []byte("1234567890123456"), //默认16位
	}
}

// SetSign 设置签名
func (s *JwtAuth) SetSign(sign string) Methods {
	if len(sign) == 0 {
		log.Println("err ：sign is '' , must has len")
	}

	s.sign = []byte(sign)
	return s
}

// SetExpire 设置过期时间
func (s *JwtAuth) SetExpire(expire time.Duration) Methods {
	if expire <= 0 {
		return s
	}
	s.expire = expire
	return s
}

type Claims struct {
	Data []byte
	jwt.StandardClaims
}

// Encrypt 加密
func (s *JwtAuth) Encrypt(rawData any) (str string, err error) {
	by, err := json.Marshal(rawData)
	if err != nil {
		return
	}
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.expire).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    string(s.sign),
			Subject:   "user token",
		},
		Data: by,
	}

	str, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.sign)
	if err != nil {
		return
	}
	return
}

// Decrypt 解密
func (s *JwtAuth) Decrypt(str string, decryptData any) (data EncryptData, err error) {
	_, claims, err := s.ParseToken(str)
	if err != nil {
		return
	}
	err = json.Unmarshal(claims.Data, &decryptData)
	if err != nil {
		return
	}
	data.Expire = claims.ExpiresAt
	data.Data = claims.Data
	return
}

// DecryptExpire 解密并判断是否过期
func (s *JwtAuth) DecryptExpire(str string, decryptData any) (data EncryptData, expire bool, err error) {
	token, claims, err := s.ParseToken(str)
	if err != nil {
		return
	}
	err = json.Unmarshal(claims.Data, &decryptData)
	if err != nil {
		return
	}
	data.Expire = claims.ExpiresAt
	data.Data = claims.Data
	return data, token.Valid, err
}
func (s *JwtAuth) ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return s.sign, nil
	})
	return token, claims, err
}
