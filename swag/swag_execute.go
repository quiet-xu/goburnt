package swag

import (
	"fmt"
	"strings"
)

// getExecuteData 数据处理，每段注释变成可读的内容
func (s SwagRead) getExecuteData(write string) (base ReadSwagBase, err error) {
	firstConv := strings.ReplaceAll(write, "\n", "")
	outList := strings.Split(firstConv, "@")
	for _, item := range outList {
		if strings.Index(item, "Router") >= 0 {
			route := s.getRouter(item)
			base.Router = route.Url
			base.Method = route.Method
		}
		if strings.Index(item, "Summary") >= 0 {
			base.Name = s.getSummary(item)
		}
		if strings.Index(item, "Authorization") >= 0 {
			base.Auth = true
		}
		if strings.Index(item, "Description") >= 0 {
			base.Description += fmt.Sprintf("%s\n", s.getDescription(item))
		}
	}
	return
}

type RouteItem struct {
	Url    string
	Method string
}

// getRouter 获取路由
func (SwagRead) getRouter(dst string) (router RouteItem) {
	i := 0
	routerStrs := strings.Split(dst, " ")
	for _, item := range routerStrs {
		if len(strings.ReplaceAll(item, " ", "")) > 0 {
			i++
			switch i {
			case 2:
				router.Url = item
			case 3:
				router.Method = item
			}
		}
	}
	return
}

// getSummary  获取标题
func (SwagRead) getSummary(dst string) (summary string) {
	i := 0
	routerStrs := strings.Split(dst, " ")
	for _, item := range routerStrs {
		if len(strings.ReplaceAll(item, " ", "")) > 0 {
			i++
			switch i {
			case 2:
				summary = item
			}
		}
	}
	return
}

// getDescription 获取详细信息
func (SwagRead) getDescription(dst string) (description string) {
	i := 0
	routerStrs := strings.Split(dst, " ")
	for _, item := range routerStrs {
		if len(strings.ReplaceAll(item, " ", "")) > 0 {
			i++
			switch i {
			case 2:
				description = item
			}
		}
	}
	return
}

//func (SwagRead)
