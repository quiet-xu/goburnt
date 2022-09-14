package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quiet-xu/goburnt/http/resp"
	"reflect"
	"strings"
)

var MiddlewareInterfaceType = reflect.TypeOf((*MiddlewareInterface)(nil)).Elem()

type MiddlewareInterface interface {
	ThisIsMiddlewareInterface()
	SetField(map[string]interface{})
}

// 公共错误类型
var (
	ReqValidateErr = fmt.Errorf("参数验证失败！")
)

type Cbt struct {
	globalErrCode  string
	response       any
	responseConfig struct {
		successFieldName string
		dataFieldName    string
		errFieldName     string
		codeFieldName    string
	}
}

func NewCbt() *Cbt {
	return &Cbt{
		response: resp.Response{},
	}
}

//SetResponse 设置自定义返回结构
/*
type Response struct {
	Success bool        `cbt:"success"`
	Data    interface{} `cbt:"data"`
	Err     FailData      `cbt:"err"`
}
支持自定义json,只需要实现 cbt tag即可
*/
func (s *Cbt) SetResponse(response any) *Cbt {
	s.response = response
	return s
}

// Cbt 转换 (基础) ControlBasicTrans
func (s Cbt) Cbt(apiFunc interface{}) func(c *gin.Context) {
	s.getResponseConfig()
	apiVal := reflect.ValueOf(apiFunc)
	if apiVal.Type().Kind() != reflect.Func {
		panic("入参需要是函数")
	}
	inputCount := apiVal.Type().NumIn()
	outputCount := apiVal.Type().NumOut()
	if inputCount != 1 {
		panic(apiVal.String() + " function only need one arg")
	}
	switch outputCount {
	case 1, 2:
		if !apiVal.Type().Out(outputCount - 1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			panic("only one value need to be error")
		}
	default:
		panic("function need return 1 or 2 value")
	}

	return func(c *gin.Context) {
		realInputRaw := reflect.New(apiVal.Type().In(0))
		realInput := realInputRaw.Interface()
		if c.Request.Method == "GET" {
			err := c.ShouldBind(realInput)
			if err != nil && err.Error() != "EOF" {
				s.FailWithData(ReqValidateErr, c)
				return
			}
		} else {
			if strings.Index(c.Request.Header.Get("Content-Type"), "multipart/form-data") < 0 {
				err := c.ShouldBindJSON(realInput)
				if err != nil && err.Error() != "EOF" {
					s.FailWithData(ReqValidateErr, c)
					return
				}
			}
		}
		if reflect.TypeOf(realInput).Implements(MiddlewareInterfaceType) {
			realInput.(MiddlewareInterface).SetField(c.Keys)
		}
		in := []reflect.Value{
			reflect.ValueOf(realInput).Elem(),
		}
		o := apiVal.Call(in)

		switch len(o) {
		case 1: //(err error)
			if o[0].Interface() != nil && o[0].Interface().(error) != nil {
				s.FailWithData(o[0].Interface().(error), c)
				return
			} else {
				s.SuccessWithData(nil, c)
			}
		case 2: //(res interface{},err error)
			if o[1].Interface() != nil && o[1].Interface().(error) != nil {
				s.FailWithData(o[1].Interface().(error), c)
				return
			} else {
				s.SuccessWithData(o[0].Interface(), c)
			}
		}
	}
}
