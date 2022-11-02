package auth

import "time"

// Methods 鉴权
type Methods interface {
	SetSign(sign string) Methods
	SetExpire(expire time.Duration) Methods
	Encrypt(rawData any) (str string, err error)
	Decrypt(str string, decryptData any) (data EncryptData, err error)
	DecryptExpire(rawData string, decryptData any) (data EncryptData, expire bool, err error)
}

type EncryptData struct {
	Data   []byte
	Expire int64
}
