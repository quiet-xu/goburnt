package burnt

import (
	"github.com/quiet-xu/goburnt/swag"
)

func (s *Burnt) setDevBoot() (err error) {
	services, err := swag.NewSwagClient(s.services...).ReadSwag()
	if err != nil {
		return
	}
	swag.NewSwagClient().PutOut(services...)
	for _, service := range services {
		for _, item := range service.Routers {
			var mids []string
			for mid := range service.Mids {
				if _, have := service.ExcludeMids[mid]; !have {
					mids = append(mids, mid)
				}
			}
			s.http.AnyByType(item.Url, service.Func, item.Method, mids...)
		}

	}
	err = s.http.Init()
	if err != nil {
		panic(err)
	}
	return
}
