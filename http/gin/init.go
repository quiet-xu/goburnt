package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quiet-xu/goburnt/http"
	"log"
	"reflect"
	"time"
)

type HttpGin struct {
	basePath  string
	port      string
	ginEngine *gin.Engine
	ginRoute  *gin.RouterGroup
	cbt       *Cbt
}

func NewClient(basePath string) *HttpGin {
	ginDefault := gin.Default()
	return &HttpGin{
		ginEngine: ginDefault,
		port:      "0.0.0.0:8080",
		basePath:  basePath,
		ginRoute:  ginDefault.Group(basePath),
		cbt:       NewCbt(),
	}
}

func (s HttpGin) Init() (err error) {

	//跨域
	s.ginEngine.Use(cors.New(cors.Config{
		AllowOriginFunc: func(string) bool {
			return true
		},
		AllowAllOrigins: false,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Length", "Content-Type", "Content-code", "Content-data"},
		MaxAge:          12 * time.Hour,
	}))
	s.ginEngine.Group(s.basePath)
	err = s.ginEngine.Run(s.port)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

//
//func (s HttpGin) SetCorsHeaders(headers ...string) http.HttpMethods {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s HttpGin) SetCorsMethods(methods ...string) http.HttpMethods {
//	//TODO implement me
//	panic("implement me")
//}
//

func (s *HttpGin) Post(api string, fv reflect.Value) http.HttpMethods {
	s.ginRoute.POST(api, s.cbt.Cbt(fv.Interface()))
	return s
}
func (s *HttpGin) Get(api string, fv reflect.Value) http.HttpMethods {
	s.ginRoute.GET(api, s.cbt.Cbt(fv.Interface()))
	return s
}

func (s *HttpGin) SetPort(port string) http.HttpMethods {
	s.port = port
	return s
}

func (s *HttpGin) SetResponse(response any) http.HttpMethods {
	s.cbt.SetResponse(response)
	return s
}

func (s *HttpGin) AnyByType(api string, fv reflect.Value, tp string) http.HttpMethods {
	switch tp {
	case "[POST]":
		s.Post(api, fv)
	case "[Get]":
		s.Get(api, fv)
	}
	return s
}
