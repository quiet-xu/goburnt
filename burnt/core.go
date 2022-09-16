package burnt

import (
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/http"
	"github.com/quiet-xu/goburnt/swag"
	"strings"
)

type Burnt struct {
	baseConf     conf.BaseConf
	http         http.HttpMethods
	services     []any
	productBurnt map[string]swag.ReadSwagBase
}

func NewBurntBuilder(services ...any) *Burnt {
	return &Burnt{
		baseConf:     *conf.DefaultBaseConf(),
		productBurnt: make(map[string]swag.ReadSwagBase),
		services:     services,
	}
}

// SetBaseConfByPath 设置默认配置 从文件读取
func (s *Burnt) SetBaseConfByPath(path string) *Burnt {
	var readConfClient conf.ReadMethods
	paths := strings.Split(path, ".")
	if len(paths) < 1 {
		panic("Unknown file type")
	}
	switch paths[1] {
	case "yaml":
		readConfClient = conf.NewReadYaml()
	default:
		panic("Unsupported file type")
	}
	s.baseConf = readConfClient.GetByPath(path)
	return s
}

// SetBaseConf 设置基础配置
func (s *Burnt) SetBaseConf(baseConf *conf.BaseConf) *Burnt {
	s.baseConf = *baseConf
	return s
}

// SetHttpConf 设置http基础配置
func (s *Burnt) SetHttpConf(methods http.HttpMethods) *Burnt {
	s.http = methods
	return s
}

// Boot 一键启动
func (s *Burnt) Boot() (err error) {
	if s.http == nil {
		s.SetDefaultHttp()
	}
	if len(s.http.GetBasePath()) == 0 {
		s.http.SetBasePath(s.baseConf.Server.Base)
	}
	if len(s.http.GetPort()) == 0 {
		s.http.SetPort(s.baseConf.Server.Port)
	}
	if s.baseConf.Server.Debug {
		err = s.setDevBoot()
		if err != nil {
			return
		}

	} else {
		err = s.setProductBoot()
		if err != nil {
			return
		}
	}

	return
}
