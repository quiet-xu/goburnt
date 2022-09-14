# github.com/quiet-xu/goburnt


'面相swagger编程'，更快的生产业务，遵守'约定大于配置'，开箱即用



swagger = > gin = > cbt => services



## Installation

To install github.com/quiet-xu/goburnt package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.18+ is required**), then you can use the below Go command to install Gin.

```sh
$ go get -u github.com/quiet-xu/github.com/quiet-xu/goburnt
```

2. Import it in your code:

```go
import "github.com/quiet-xu/github.com/quiet-xu/goburnt"
```

## Quick start DEV

```go
package main

import (
	"github.com/quiet-xu/goburnt/burnt"
	"github.com/quiet-xu/goburnt/demo"
	"github.com/quiet-xu/goburnt/conf"
)

type Services struct {
    
}

// Get 获取一个信息
// @Summary 获取一个信息（标题）
// @Description 注释1
// @Description 注释2
// @Tags 分组
// @Param Authorization header string true "身份加密串"
// @Router /a [POST]
func (s Services) Get(a string) (string,error) {
	return a, nil
}


func main() {
	burnt.NewBurntBuilder(Services{}).
		SetBaseConf(conf.DefaultBaseConf().
			SetBase("basepath").
			SetDev()).
		Boot()
}
```

## Quick start Product

```go
package main

import (
	"github.com/quiet-xu/goburnt/burnt"
	"github.com/quiet-xu/goburnt/demo"
	"github.com/quiet-xu/goburnt/conf"
)
type Services struct {

}

// Get 获取一个信息
// @Summary 获取一个信息（标题）
// @Description 注释1
// @Description 注释2
// @Tags 分组
// @Param Authorization header string true "身份加密串"
// @Router /a [POST]
func (s Services) Get(a string) (string,error) {
	return a, nil
}


func main() {
	burnt.NewBurntBuilder(Services{}).
		SetBaseConf(conf.DefaultBaseConf().
			SetBase("basepath").
			SetProduct()).
		Boot()
}
```


## 原理

通过 go doc 和 services 获取 swag注释，并自动解析 api所对应的func，dev环境采用每次运行自动生成json，product采用每次运行读取json来运行gin

## 实现

| 功能                    | 实现情况 |
|-----------------------|---|
| 单体中间件                 |   |
| 分组建中间件                | ✅ |
| 全局context(自定义context) | ✅ |
| 自动路由                  | ✅ |
| 自定义组件                 |   |
| 鉴权中心                  |  |
