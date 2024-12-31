package swag

import (
	"encoding/json"
	"fmt"
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/doc"
	"os"
	"reflect"
	"sort"
	"strings"
)

type SwagRead struct {
	methodPathMap map[string][]Method //key: path val:methods
	services      []any
}
type Method struct {
	Name       string        //方法名
	Func       reflect.Value //方法
	StructName string        //总方法名
}
type ReadSwagBase struct {
	Routers     []RouteItem    `json:"routers"`     //api
	Auth        bool           `json:"auth"`        //权限
	Name        string         `json:"name"`        //api名称
	Description string         `json:"description"` //详细介绍
	Tag         string         `json:"tag"`         //分组
	Mids        map[string]int `json:"mid"`         //中间件
	ExcludeMids map[string]struct{}
	Func        reflect.Value `json:"-"`
	StructName  string        `json:"structName"`
	PkgPath     string        `json:"pkgPath"`
	FuncName    string        `json:"funcName"`
}

func NewSwagClient(services ...any) *SwagRead {
	return &SwagRead{
		services:      services,
		methodPathMap: make(map[string][]Method),
	}
}

// ReadSwag 读取swag注释
func (s *SwagRead) ReadSwag() (bases []ReadSwagBase, err error) {
	for _, service := range s.services {
		s.getMethodNameAndPkgPaths(service)
	}
	for key, methods := range s.methodPathMap {
		for _, method := range methods {
			var out string
			out, err = doc.NewCmdClient().ReadDocByStructAndMethodName(key, method.StructName, method.Name)
			if err != nil {
				return
			}
			if len(out) == 0 {
				continue
			}
			var base ReadSwagBase

			base, err = s.executeServicesData(out)
			if err != nil {
				return
			}
			base.FuncName = method.Name
			base.Func = method.Func
			base.PkgPath = key
			base.StructName = method.StructName
			out, err = doc.NewCmdClient().ReadDocByMethodName(key, method.StructName)
			if err != nil {
				return
			}
			if len(out) == 0 {
				continue
			}
			groupMids := s.executeMethodGroupData(out).Mids
			for k, val := range groupMids {
				if base.Mids == nil {
					base.Mids = make(map[string]int)
				}
				base.Mids[k] = val
			}
			bases = append(bases, base)
		}
	}
	if len(bases) > 0 {
		var buf []byte
		buf, err = json.Marshal(bases)
		if err != nil {
			return
		}
		var file *os.File
		path := "burnt.json"
		if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".json") {
			path = os.Args[1]
		}
		path = conf.LocalFileAuto(path)
		file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		_, err = file.Write(buf)
		if err != nil {
			return
		}
	}

	return
}

// AddReadServices 添加服务
func (s *SwagRead) AddReadServices(services ...any) *SwagRead {
	s.services = append(s.services, services)
	return s
}

// getMethodNameAndPkgPaths 获取方法名 + pkgPath
func (s *SwagRead) getMethodNameAndPkgPaths(dst any) {
	stValue := reflect.ValueOf(dst)
	fmt.Println("nn", stValue.NumMethod())
	if stValue.Kind() == reflect.Ptr {
		stValue = stValue.Elem()
	}
	dtype := reflect.TypeOf(dst)
	if dtype.Kind() == reflect.Ptr {
		dtype = dtype.Elem()
	}
	structName := dtype.Name()
	num := reflect.ValueOf(dst).NumMethod()
	for i := 0; i < num; i++ {
		name := reflect.TypeOf(dst).Method(i).Name
		path := dtype.PkgPath()
		s.methodPathMap[path] = append(s.methodPathMap[path], Method{
			Name:       name,
			Func:       reflect.ValueOf(dst).Method(i),
			StructName: structName,
		})
	}
	return
}

// PutOut 输出已经读到的api
func (s *SwagRead) PutOut(base ...ReadSwagBase) {
	fmt.Println(fmt.Sprintf("[序号]  %-10s %-60s %-25s %-60s %v", "方法", "api", "分组", "描述", "中间件"))
	for i, swagBase := range base {
		for k, route := range swagBase.Routers {
			var midIndexS []int
			midNewMap := make(map[int]string)
			for mid, v := range swagBase.Mids {
				if _, have := swagBase.ExcludeMids[mid]; !have {
					midIndexS = append(midIndexS, v)
					midNewMap[v] = mid
				}
			}
			sort.Ints(midIndexS)
			var mids []string
			for _, index := range midIndexS {
				mids = append(mids, midNewMap[index])
			}
			fmt.Println(fmt.Sprintf("[%-5s]%-10s %-60s %-25s %-60s %v", fmt.Sprintf("%d-%d", i+1, k), route.Method, route.Url, swagBase.Tag, swagBase.Name, mids))
		}
	}
}
