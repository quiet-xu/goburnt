package ghttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func (s *Cbt) getResponseConfig() {
	tp := reflect.TypeOf(s.response)
	num := tp.NumField()
	for i := 0; i < num; i++ {
		cbt := tp.Field(i).Tag.Get("cbt")
		json := tp.Field(i).Tag.Get("json")
		switch cbt {
		case "success":
			if len(json) > 0 {
				s.responseConfig.successFieldName = json
			} else {
				s.responseConfig.successFieldName = tp.Field(i).Name
			}
		case "data":
			if len(json) > 0 {
				s.responseConfig.dataFieldName = json
			} else {
				s.responseConfig.dataFieldName = tp.Field(i).Name
			}
		case "err":
			if len(json) > 0 {
				s.responseConfig.errFieldName = json
			} else {
				s.responseConfig.errFieldName = tp.Field(i).Name
			}
		case "code":
			if len(json) > 0 {
				s.responseConfig.codeFieldName = json
			} else {
				s.responseConfig.codeFieldName = tp.Field(i).Name
			}
		}

	}
}

// SuccessWithData 成功参数
func (s Cbt) SuccessWithData(data interface{}, c *gin.Context) {
	resp := make(gin.H)
	if len(s.responseConfig.dataFieldName) > 0 {
		resp[s.responseConfig.dataFieldName] = data
	}
	if len(s.responseConfig.successFieldName) > 0 {
		resp[s.responseConfig.successFieldName] = true
	}
	c.JSON(http.StatusOK, resp)
}

// FailWithData 失败参数
func (s Cbt) FailWithData(err error, c *gin.Context) {
	resp := make(gin.H)
	if len(s.responseConfig.errFieldName) > 0 {
		resp[s.responseConfig.errFieldName] = err
	} else if len(s.responseConfig.dataFieldName) > 0 {
		resp[s.responseConfig.dataFieldName] = err
	}
	if len(s.responseConfig.successFieldName) > 0 {
		resp[s.responseConfig.successFieldName] = false
	}
	c.JSON(http.StatusOK, resp)
	c.Abort()
}

// FailWithCode 失败参数
func (s Cbt) FailWithCode(err string, code int, c *gin.Context) {
	resp := make(gin.H)
	if len(s.responseConfig.errFieldName) > 0 {
		resp[s.responseConfig.errFieldName] = err
	} else if len(s.responseConfig.dataFieldName) > 0 {
		resp[s.responseConfig.dataFieldName] = err
	}
	if len(s.responseConfig.successFieldName) > 0 {
		resp[s.responseConfig.successFieldName] = false
	}
	c.JSON(code, resp)
	c.Abort()
}
