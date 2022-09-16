package burnt

import (
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/http/gin"
)

func GetHttpGinDefault() *gin.HttpGin {
	return gin.GetGinConf()
}

func GetConfDefault() *conf.BaseConf {
	return conf.DefaultBaseConf()
}

func (s *Burnt) SetDefaultHttp() {
	s.http = GetHttpGinDefault().
		SetBasePath(s.baseConf.Server.Base).
		SetPort(s.baseConf.Server.Port)
}
