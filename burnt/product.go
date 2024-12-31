package burnt

import (
	"encoding/json"
	"fmt"
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
		t := reflect.TypeOf(view)
		if stValue.Kind() == reflect.Ptr {
			stValue = stValue.Elem()
			t = t.Elem()
		}
		structName := t.Name()
		fmt.Println(structName)
		num := reflect.ValueOf(view).NumMethod()
		for i := 0; i < num; i++ {
			name := reflect.TypeOf(view).Method(i).Name
			path := t.PkgPath()
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
