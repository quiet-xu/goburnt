package http

import "reflect"

type HttpMethods interface {
	//// SetCorsHeaders cors allowHeaders default "Origin", "Authorization", "Content-Length", "Content-Type", "Content-code"
	//SetCorsHeaders(headers ...string) HttpMethods //
	//
	//// SetCorsMethods GET POST PUT DELETE
	//SetCorsMethods(methods ...string) HttpMethods

	AnyByType(api string, fv reflect.Value, tp string) HttpMethods

	Get(api string, fv reflect.Value) HttpMethods

	Post(api string, fv reflect.Value) HttpMethods

	SetPort(port string) HttpMethods

	SetResponse(response any) HttpMethods

	Init() error
	//
}
