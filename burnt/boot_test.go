package burnt

import (
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/demo"
	"testing"
)

func TestName(t *testing.T) {
	err := NewBurntBuilder(demo.FView{}).
		SetBaseConf(conf.DefaultBaseConf().
			SetBase("asdasd").
			SetProduct()).
		Boot()
	t.Log(err)
}
