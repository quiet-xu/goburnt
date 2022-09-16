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

	for _, service := range s.services {
		stValue := reflect.ValueOf(service)
		structName := reflect.TypeOf(service).Name()
		num := stValue.NumMethod()
		for i := 0; i < num; i++ {
			name := reflect.TypeOf(service).Method(i).Name
			path := reflect.TypeOf(service).PkgPath()
			if val, has := s.productBurnt[path+structName+name]; has {
				s.http.AnyByType(val.Router, reflect.ValueOf(service).Method(i), val.Method, val.Mids...)
			}
		}
	}

	err = s.http.Init()
	if err != nil {
		panic(err)
	}
	return
}
