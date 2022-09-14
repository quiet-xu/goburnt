package swag

import (
	"testing"
)

type X struct {
	A string
	B string
}

func TestName(t *testing.T) {
	x := []X{
		{A: "13", B: "caca"},
		{A: "13", B: "vb"},
		{A: "13", B: "caca"},
		{A: "24", B: "vb"},
		{A: "24", B: "caca"},
		{A: "24", B: "caca"},
	}

	a := make(map[int]*X)
	for index, x2 := range x {
		c := x2
		a[index] = &c
	}
	t.Log(a)
	//a, b := NewSwagClient(demo.FView{}).ReadSwag()
	//t.Log(a, b)

	//SwagRead{}.GetMethodNames(demo.FView{})
}
