package conf

type BaseConf struct {
	Server Server
	Db     map[string]Database
	Redis  map[string]Redis
}
type Database struct {
	Url   string
	Debug bool
	Cache bool
}
type Redis struct {
	Prefix string
	Driver string
	Url    string
	Pass   string
	Db     int
	Expire int
}
type Server struct {
	Port  string
	Debug bool
	Sk    string
	Base  string
}

var baseConf *BaseConf

func DefaultBaseConf() *BaseConf {
	baseConf = &BaseConf{
		Server: Server{
			Base:  "/base",
			Port:  "0.0.0.0:8080",
			Debug: true,
			Sk:    "1234567890123456",
		},
	}
	return baseConf
}

// SetPort default 8080
func (s *BaseConf) SetPort(port string) *BaseConf {
	s.Server.Port = port
	return s
}

// SetProduct 生产环境
func (s *BaseConf) SetProduct() *BaseConf {
	s.Server.Debug = false
	return s
}

// SetDev 测试环境
func (s *BaseConf) SetDev() *BaseConf {
	s.Server.Debug = true
	return s
}

// SetSk 秘钥
func (s *BaseConf) SetSk(sk string) *BaseConf {
	s.Server.Sk = sk
	return s
}

// SetBase base 路径
func (s *BaseConf) SetBase(base string) *BaseConf {
	s.Server.Base = base
	return s
}

// GetBase 获取配置
func (BaseConf) GetBase() BaseConf {
	return *baseConf
}
