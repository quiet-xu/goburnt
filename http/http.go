package http

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type HttpMethods interface {
	//// SetCorsHeaders cors allowHeaders default "Origin", "Authorization", "Content-Length", "Content-Type", "Content-code"
	//SetCorsHeaders(headers ...string) HttpMethods //
	//
	//// SetCorsMethods GET POST PUT DELETE
	//SetCorsMethods(methods ...string) HttpMethods

	AnyByType(path string, fv reflect.Value, tp string, mids ...string) HttpMethods

	Get(path string, fv reflect.Value, mids ...string) HttpMethods

	Post(path string, fv reflect.Value, mids ...string) HttpMethods

	SetPort(port string) HttpMethods
	GetPort() string

	SetResponse(response any) HttpMethods

	SetBasePath(basePath string) HttpMethods
	GetBasePath() string

	SetMidFunc(midName string, mid func(*gin.Context)) HttpMethods

	Init() error

	SetLoadHtml(pattern string) HttpMethods
	//
}
