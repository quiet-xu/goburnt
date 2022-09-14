package burnt

import (
	"github.com/quiet-xu/goburnt/conf"
	"github.com/quiet-xu/goburnt/http/resp"
	"github.com/quiet-xu/goburnt/swag"
	"strings"
)

type Burnt struct {
	baseConf     conf.BaseConf
	services     []any
	productBurnt map[string]swag.ReadSwagBase
	response     any
}

func NewBurntBuilder(services ...any) *Burnt {
	return &Burnt{
		baseConf:     *conf.DefaultBaseConf(),
		productBurnt: make(map[string]swag.ReadSwagBase),
		services:     services,
		response:     resp.Response{},
	}
}
func (s *Burnt) SetResponse(response any) *Burnt {
	s.response = response
	return s
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
func (s Burnt) Boot() (err error) {
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
