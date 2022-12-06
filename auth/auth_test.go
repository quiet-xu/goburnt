package auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type X struct {
	A string
}

func TestAsc(t *testing.T) {
	var x X
	x.A = "123"
	client := NewAscAuth().SetSign("SJX")
	str, _ := client.Encrypt(x)
	t.Log(str)
	var x1 X
	_, ok, _ := client.DecryptExpire(str, &x1)
	t.Log(x1, ok)
}
func TestJwt(t *testing.T) {
	var x X
	x.A = "123"
	client := NewJwtAuth().SetSign("token")
	str, _ := client.Encrypt(x)
	t.Log(str)
	var x1 X
	_, ok, _ := client.DecryptExpire(str, &x1)
	t.Log(x1, ok)
}
func TestDecodeJwt(t *testing.T) {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzM4MSwidXNlcm5hbWUiOiLmtYvor5Xlj7ciLCJhZG1pbiI6MSwiaWF0IjoxNjcwMjM3MTczLCJleHAiOjE2NzAyODAzNzN9.fqDvAnRCzAS9iZvtGa3oDSMhtEBf15_858xYmpfMBrc"
	claims := Claims{}
	//jwt.ParseWithClaims()
	token, err := jwt.ParseWithClaims(str, claims, func(token *jwt.Token) (i interface{}, err error) {
		return "token", nil
	})

	if err != nil {
		return
	}
	t.Log(token)
}
