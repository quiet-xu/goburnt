package burnt

import (
	"encoding/json"
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/swag"
	"reflect"
	"sort"
)

func (s *Burnt) setProductBoot() (err error) {
	buf, err := conf.NewReadJson().GetBytesByPath("burnt.json")
	var services []swag.ReadSwagBase
	err = json.Unmarshal(buf, &services)
	if err != nil {
		return
	}
	swag.NewSwagClient().PutOut(services...)
	for _, service := range services {
		s.productBurnt[service.PkgPath+service.StructName+service.FuncName] = service
	}
	for _, view := range s.services {
		stValue := reflect.ValueOf(view)
		structName := reflect.TypeOf(view).Name()
		num := stValue.NumMethod()
		for i := 0; i < num; i++ {
			name := reflect.TypeOf(view).Method(i).Name
			path := reflect.TypeOf(view).PkgPath()
			if val, has := s.productBurnt[path+structName+name]; has {
				var midIndexS []int
				midNewMap := make(map[int]string)
				for mid, v := range val.Mids {
					if _, have := val.ExcludeMids[mid]; !have {
						midIndexS = append(midIndexS, v)
						midNewMap[v] = mid
					}
				}
				sort.Ints(midIndexS)
				var mids []string
				for _, index := range midIndexS {
					mids = append(mids, midNewMap[index])
				}
				for _, item := range val.Routers {
					s.http.AnyByType(item.Url, reflect.ValueOf(view).Method(i), item.Method, mids...)
				}
			}

		}
	}

	err = s.http.Init()
	if err != nil {
		panic(err)
	}
	return
}
