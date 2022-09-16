package burnt

import (
	"github.com/quiet-xu/goburnt/demo"
	"github.com/quiet-xu/goburnt/demo/mid"
	"testing"
)

func TestName(t *testing.T) {
	err := NewBurntBuilder(demo.FView{}).
		SetBaseConf(GetConfDefault().
			SetBase("asdasd").
			SetDev()).
		SetHttpConf(GetHttpGinDefault().
			SetMidFunc("auth", mid.Auth).
			SetMidFunc("log", mid.Log),
		).
		Boot()
	t.Log(err)
}

func TestNameByPath(t *testing.T) {
	err := NewBurntBuilder(demo.FView{}).
		SetBaseConfByPath("demo/config.yaml").
		Boot()
	t.Log(err)
}
