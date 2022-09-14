package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type ReadMethods interface {
	GetByPath(path string) BaseConf //根据路径看配置
	GetBytesByPath(path string) ([]byte, error)
}
type ReadYaml struct {
}

func NewReadYaml() *ReadYaml {
	return &ReadYaml{}
}

// GetByPath 根据路径获取配置
func (ReadYaml) GetByPath(path string) BaseConf {
	if len(path) == 0 {
		path = "config.yaml"
	}
	if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".yaml") {
		path = os.Args[1]
	}
	path = LocalFileAuto(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(buf, baseConf); nil != err {
		panic(err)
	}
	return *baseConf
}

// GetBytesByPath 根据路径获取byte类型的配置
func (ReadYaml) GetBytesByPath(path string) (buf []byte, err error) {
	if len(path) == 0 {
		path = "config.yaml"
	}
	if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".yaml") {
		path = os.Args[1]
	}
	path = LocalFileAuto(path)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	buf, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return
}
