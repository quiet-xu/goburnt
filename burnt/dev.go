package burnt

import (
	"github.com/quiet-xu/goburnt/swag"
)

func (s Burnt) setDevBoot() (err error) {
	services, err := swag.NewSwagClient(s.services...).ReadSwag()
	if err != nil {
		return
	}
	for _, service := range services {
		s.http.AnyByType(service.Router, service.Func, service.Method, service.Mids...)
	}
	err = s.http.Init()
	if err != nil {
		panic(err)
	}
	return
}
