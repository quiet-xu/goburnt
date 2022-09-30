package auth

import "testing"

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
	client := NewJwtAuth().SetSign("SJX")
	str, _ := client.Encrypt(x)
	t.Log(str)
	var x1 X
	_, ok, _ := client.DecryptExpire(str, &x1)
	t.Log(x1, ok)
}
