package swag

import (
	"fmt"
	"strings"
)

// executeServicesData 数据处理，每段注释变成可读的内容
func (s *SwagRead) executeServicesData(write string) (base ReadSwagBase, err error) {
	firstConv := strings.ReplaceAll(write, "\n", "")
	outList := strings.Split(firstConv, "@")
	if len(outList) > 1 {
		outList = outList[1:]
	} else {
		return
	}

	for _, item := range outList {
		if strings.Contains(item, "Success") || strings.Contains(item, "Param") {
			continue
		}
		if strings.Contains(item, "Router") {
			route := s.getRouter(item)
			base.Routers = append(base.Routers, RouteItem{
				Url:    route.Url,
				Method: route.Method,
			})
		}
		if strings.Contains(item, "Tag") {
			base.Tag = s.getTag(item)
		}

		if strings.Contains(item, "Summary") {
			base.Name = s.getSummary(item)
		}
		if strings.Contains(item, "Authorization") ||
			strings.Contains(item, "Token") {
			base.Auth = true
		}
		if strings.Contains(item, "Mid") && !strings.Contains(item, "Mid!") {
			base.mids = append(base.mids, s.getMid(item, 2)...)
		}
		if strings.Contains(item, "Mid!") {
			base.ExcludeMids = s.getExcludeMid(item)
		}
		if strings.Contains(item, "Description") {
			base.Description += fmt.Sprintf("%s\n", s.getDescription(item))
		}
	}
	return
}

// executeMethodGroupData 处理方法组的注释
func (s *SwagRead) executeMethodGroupData(write string) (mids []string) {
	outList := strings.Split(write, "\n")
	for _, outItem := range outList {
		if strings.Contains(outItem, "@Mid") {
			outItem = strings.TrimPrefix(outItem, " ")
			mids = append(mids, s.getMid(outItem, 3)...)
		}
	}
	return
}

type RouteItem struct {
	Url    string
	Method string
}

// getRouter 获取路由
func (*SwagRead) getRouter(dst string) (router RouteItem) {
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
func (*SwagRead) getSummary(dst string) (summary string) {
	summary = strings.ReplaceAll(dst, "Summary ", "")
	return
}

// getTag  获取标题
func (*SwagRead) getTag(dst string) (tag string) {
	tag = strings.ReplaceAll(dst, "Tag ", "")
	return
}

// getDescription 获取详细信息
func (*SwagRead) getDescription(dst string) (description string) {
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

func (*SwagRead) getMid(dst string, index int) (mids []string) {
	i := 0
	routerStrs := strings.Split(dst, " ")
	for _, item := range routerStrs {
		if len(strings.ReplaceAll(item, " ", "")) > 0 {
			i++
			switch i {
			case index:
				mids = strings.Split(item, ",")
			}
		}
	}
	return
}

func (*SwagRead) getExcludeMid(dst string) (exMap map[string]struct{}) {

	i := 0
	routerStrs := strings.Split(dst, " ")
	var mids []string
	for _, item := range routerStrs {
		if len(strings.ReplaceAll(item, " ", "")) > 0 {
			i++
			switch i {
			case 2:
				mids = strings.Split(item, ",")
			}
		}
	}
	exMap = make(map[string]struct{}, len(mids))
	for _, mid := range mids {
		exMap[mid] = struct{}{}
	}
	return
}

//func (SwagRead)
