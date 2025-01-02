package ghttp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quiet-xu/goburnt/http/resp"
	"net/http"
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
	s := &Cbt{
		response: resp.Response{},
	}
	s.getResponseConfig()
	return s
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
	s.getResponseConfig()
	return s
}

// Cbt 转换 (基础) ControlBasicTrans
func (s *Cbt) Cbt(apiFunc interface{}) func(c *gin.Context) {

	switch apiFunc.(type) {
	case func(*gin.Context):
		return apiFunc.(func(*gin.Context))
	}
	apiVal := reflect.ValueOf(apiFunc)
	if apiVal.Type().Kind() != reflect.Func {
		panic("入参需要是函数")
	}
	inputCount := apiVal.Type().NumIn()
	outputCount := apiVal.Type().NumOut()
	if inputCount < 1 {
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
		if apiVal.Type().In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
			panic("第一个参数必须是context")
		}
		in := []reflect.Value{
			reflect.ValueOf(c.Request.Context()),
		}
		if inputCount > 1 {
			realInput := reflect.New(apiVal.Type().In(1)).Interface()
			if c.Request.Method == "GET" {
				err := c.ShouldBindQuery(realInput)
				if err != nil && err.Error() != "EOF" {
					s.FailWithData(ReqValidateErr, c)
					return
				}
				err = c.ShouldBind(realInput)
				if err != nil && err.Error() != "EOF" {
					s.FailWithData(ReqValidateErr, c)
					return
				}
			} else {
				if strings.Index(c.Request.Header.Get("Content-Type"), "multipart/form-data") < 0 {
					switch realInput.(type) {
					case *[]uint8:
						buf := make([]byte, c.Request.ContentLength)
						_, err := c.Request.Body.Read(buf)
						if err != nil && err.Error() != "EOF" {
							s.FailWithData(ReqValidateErr, c)
							return
						}
						realInput = &buf
						//realInput =
						//reflect.ValueOf(realInput).Set(reflect.ValueOf(reqBytes))
						//realInput = reqBytes
					default:
						err := c.ShouldBindJSON(realInput)
						if err != nil && err.Error() != "EOF" {
							s.FailWithData(ReqValidateErr, c)
							return
						}
					}

				}
			}

			if reflect.TypeOf(realInput).Implements(MiddlewareInterfaceType) {
				realInput.(MiddlewareInterface).SetField(c.Keys)
			}
			in = append(in, reflect.ValueOf(realInput).Elem())
		}
		o := apiVal.Call(in)
		switch len(o) {
		case 1: //(err error)
			if o[0].Interface() != nil {
				switch o[0].Interface().(type) {
				case *[]uint8:
					_, err := c.Writer.Write(o[0].Interface().([]byte))
					if err != nil {
						c.Status(http.StatusBadRequest)
						return
					}
				case error:
					s.FailWithData(o[0].Interface().(error), c)
					return
				default:
					s.SuccessWithData(o[0].Interface(), c)
				}
			} else {
				s.SuccessWithData(nil, c)
			}
		case 2: //(res interface{},err error)

			if o[1].Interface() != nil && o[1].Interface().(error) != nil {
				switch o[1].Interface().(type) {
				case []uint8:
					_, err := c.Writer.Write(o[1].Interface().([]byte))
					if err != nil {
						c.Status(http.StatusBadRequest)
						return
					}
				case error:
					s.FailWithData(o[1].Interface().(error), c)
				default:
					s.SuccessWithData(o[0].Interface(), c)
					return
				}
			} else {
				switch o[0].Interface().(type) {
				case []uint8:
					_, err := c.Writer.Write(o[0].Interface().([]byte))
					if err != nil {
						c.Status(http.StatusBadRequest)
						return
					}
				case error:
					if o[0].Interface().(error) != nil {
						s.FailWithData(o[0].Interface().(error), c)
					}
					return
				default:
					s.SuccessWithData(o[0].Interface(), c)
				}
				//s.SuccessWithData(o[0].Interface(), c)
			}
		}
	}
}
