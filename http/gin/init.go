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
	basePath   string
	port       string
	ginEngine  *gin.Engine
	ginRoute   *gin.RouterGroup
	midFuncMap map[string]func(*gin.Context)
	cbt        *Cbt
}

func GetGinConf() *HttpGin {
	gin.SetMode(gin.ReleaseMode)
	ginDefault := gin.Default()
	return &HttpGin{
		ginEngine:  ginDefault,
		cbt:        NewCbt(),
		midFuncMap: make(map[string]func(*gin.Context)),
	}
}

func (s *HttpGin) Init() (err error) {

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

func (s *HttpGin) Post(path string, fv reflect.Value, mids ...string) http.HttpMethods {
	if s.ginRoute == nil {
		s.ginRoute = s.ginEngine.Group("")
	}
	handlers := []gin.HandlerFunc{
		s.cbt.Cbt(fv.Interface()),
	}
	for _, mid := range mids {
		if val, has := s.midFuncMap[mid]; has {
			handlers = append(handlers, val)
		}
	}
	s.ginRoute.POST(path, handlers...)

	return s
}
func (s *HttpGin) Get(path string, fv reflect.Value, mids ...string) http.HttpMethods {
	if s.ginRoute == nil {
		s.ginRoute = s.ginEngine.Group("")
	}
	handlers := []gin.HandlerFunc{
		s.cbt.Cbt(fv.Interface()),
	}
	for _, mid := range mids {
		if val, has := s.midFuncMap[mid]; has {
			handlers = append(handlers, val)
		}
	}
	s.ginRoute.GET(path, handlers...)

	return s
}

func (s *HttpGin) SetPort(port string) http.HttpMethods {
	s.port = port
	return s
}

func (s *HttpGin) SetResponse(response any) http.HttpMethods {
	if response == nil {
		return s
	}
	s.cbt.SetResponse(response)
	return s
}

// SetMidFunc 设置中间件
func (s *HttpGin) SetMidFunc(midName string, mid func(*gin.Context)) http.HttpMethods {
	s.midFuncMap[midName] = mid
	return s
}

// AnyByType 任何类型的接口
func (s *HttpGin) AnyByType(path string, fv reflect.Value, tp string, mids ...string) http.HttpMethods {
	switch tp {
	case "[Get]":
		s.Get(path, fv, mids...)
	default:
		s.Post(path, fv, mids...)
	}
	return s
}

// SetBasePath 设置基础路径
func (s *HttpGin) SetBasePath(basePath string) http.HttpMethods {
	s.ginRoute = s.ginEngine.Group(basePath)
	return s
}
func (s *HttpGin) GetBasePath() string {
	return s.basePath
}

func (s *HttpGin) GetPort() string {
	return s.port
}
