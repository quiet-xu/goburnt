package conf

import (
	"io/ioutil"
	"os"
	"strings"
)

type ReadJson struct {
}

func NewReadJson() *ReadJson {
	return &ReadJson{}
}

func (r ReadJson) GetBytesByPath(path string) (buf []byte, err error) {
	if len(path) == 0 {
		path = "burnt.json"
	}
	if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], ".json") {
		path = os.Args[1]
	}
	path = LocalFileAuto(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf, err = ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return
}
