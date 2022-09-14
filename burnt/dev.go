package burnt

import (
	"github.com/quiet-xu/goburnt/http"
	"github.com/quiet-xu/goburnt/http/gin"
	"github.com/quiet-xu/goburnt/swag"
)

func (s Burnt) setDevBoot() (err error) {
	services, err := swag.NewSwagClient(s.services...).ReadSwag()
	if err != nil {
		return
	}
	httpBoot := http.HttpMethods(gin.NewClient(s.baseConf.Server.Base))
	for _, service := range services {
		httpBoot.AnyByType(service.Router, service.Func, service.Method)
	}
	err = httpBoot.Init()
	if err != nil {
		panic(err)
	}
	return
}
