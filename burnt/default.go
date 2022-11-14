package burnt

import (
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/http/ghttp"
)

func GetHttpGinDefault() *ghttp.HttpGin {
	return ghttp.GetGinConf()
}

func GetConfDefault() *conf.BaseConf {
	return conf.DefaultBaseConf()
}

func (s *Burnt) SetDefaultHttp() {
	s.http = GetHttpGinDefault().
		SetBasePath(s.baseConf.Server.Base).
		SetPort(s.baseConf.Server.Port)
}
