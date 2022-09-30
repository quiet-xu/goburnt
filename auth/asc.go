package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/quiet-xu/goburnt/encrypt"
	"log"
	"strings"
	"time"
)

type AscAuth struct {
	sign   []byte        //默认16位
	expire time.Duration //默认 1小时
}

func NewAscAuth() *AscAuth {
	return &AscAuth{
		expire: time.Hour,
		sign:   []byte("1234567890123456"), //默认16位
	}
}

// SetSign 设置签名
func (s *AscAuth) SetSign(sign string) Methods {
	if strings.Count(sign, "") != 16 || strings.Count(sign, "") != 32 {
		log.Println("err ：sign must 16 / 32。so，sign use def")
		return s
	}
	s.sign = []byte(sign)
	return s
}

// SetExpire 设置过期时间
func (s *AscAuth) SetExpire(expire time.Duration) Methods {
	if expire <= 0 {
		return s
	}
	s.expire = expire
	return s
}

// Encrypt 加密
func (s *AscAuth) Encrypt(rawData any) (str string, err error) {
	by, err := json.Marshal(rawData)
	if err != nil {
		return
	}
	data := EncryptData{
		Data:   by,
		Expire: time.Now().Add(s.expire).Unix(),
	}
	by, err = json.Marshal(data)
	if err != nil {
		return
	}
	by, err = encrypt.AesCBCEncrypt(by, s.sign)
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(by), nil
}

// Decrypt 解密
func (s *AscAuth) Decrypt(str string, decryptData any) (data EncryptData, err error) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	by, err = encrypt.AesCBCDncrypt(by, s.sign)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, &data)
	if err != nil {
		return
	}
	err = json.Unmarshal(data.Data, &decryptData)
	if err != nil {
		return
	}
	return
}

// DecryptExpire 解密并判断是否过期
func (s *AscAuth) DecryptExpire(rawData string, decryptData any) (data EncryptData, expire bool, err error) {
	data, err = s.Decrypt(rawData, decryptData)
	if err != nil {
		return
	}
	if data.Expire < time.Now().Unix() {
		expire = false
		return
	}
	expire = true
	return
}
