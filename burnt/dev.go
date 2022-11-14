package burnt

import (
	"github.com/quiet-xu/goburnt/swag"
	"sort"
)

func (s *Burnt) setDevBoot() (err error) {
	services, err := swag.NewSwagClient(s.services...).ReadSwag()
	if err != nil {
		return
	}
	swag.NewSwagClient().PutOut(services...)
	for _, service := range services {
		var midIndexS []int
		midNewMap := make(map[int]string)
		for mid, v := range service.Mids {
			if _, have := service.ExcludeMids[mid]; !have {
				midIndexS = append(midIndexS, v)
				midNewMap[v] = mid
			}
		}
		sort.Ints(midIndexS)
		var mids []string
		for _, index := range midIndexS {
			mids = append(mids, midNewMap[index])
		}
		for _, item := range service.Routers {
			s.http.AnyByType(item.Url, service.Func, item.Method, mids...)
		}
	}
	err = s.http.Init()
	if err != nil {
		panic(err)
	}
	return
}
