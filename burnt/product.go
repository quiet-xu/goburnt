package burnt

import (
	"encoding/json"
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/swag"
	"reflect"
)

func (s *Burnt) setProductBoot() (err error) {
	buf, err := conf.NewReadJson().GetBytesByPath("burnt.json")
	var services []swag.ReadSwagBase
	err = json.Unmarshal(buf, &services)
	if err != nil {
		return
	}
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
				var mids []string
				for mid := range val.Mids {
					if _, have := val.ExcludeMids[mid]; !have {
						mids = append(mids, mid)
					}
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
