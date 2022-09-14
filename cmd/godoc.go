package cmd

import (
	"log"
	"os/exec"
)

type docType string

const (
	DocTypeC     = docType("-c")
	DocTypeAll   = docType("-all")
	DocTypeShort = docType("-short") //默认项
	DocTypeSrc   = docType("-src")
	DocTypeU     = docType("-u")
)

type Cmd struct {
	docType docType //-c -all -short -src -u
}

func NewCmdClient() *Cmd {
	return &Cmd{
		docType: DocTypeShort,
	}
}

// ReadDocByMethodName 读注释
func (s Cmd) ReadDocByMethodName(packageName, methodName string) (outStr string, err error) {
	cmd := exec.Command("go", "doc", string(s.docType), packageName, methodName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), err

}

// ReadDocByPackageName 读注释
func (s Cmd) ReadDocByPackageName(packageName string) (outStr string, err error) {
	cmd := exec.Command("go", "doc", string(s.docType), packageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), err
}

// SetDocType 阅读文档的类型
func (s *Cmd) SetDocType(docType docType) *Cmd {
	s.docType = docType
	return s
}
