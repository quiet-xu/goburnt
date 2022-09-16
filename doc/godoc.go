package doc

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

type GoDoc struct {
	docType docType //-c -all -short -src -u
}

func NewCmdClient() *GoDoc {
	return &GoDoc{
		docType: DocTypeShort,
	}
}

// ReadDocByStructAndMethodName 读注释
func (s GoDoc) ReadDocByStructAndMethodName(packageName, structName string, methodName string) (outStr string, err error) {
	cmd := exec.Command("go", "doc", string(s.docType), packageName, structName+"."+methodName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), err
}

// ReadDocByMethodName 读注释
func (s GoDoc) ReadDocByMethodName(packageName, methodName string) (outStr string, err error) {
	cmd := exec.Command("go", "doc", string(s.docType), packageName, methodName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), err

}

// ReadDocByPackageName 读注释
func (s GoDoc) ReadDocByPackageName(packageName string) (outStr string, err error) {
	cmd := exec.Command("go", "doc", string(s.docType), packageName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), err
}

// SetDocType 阅读文档的类型
func (s *GoDoc) SetDocType(docType docType) *GoDoc {
	s.docType = docType
	return s
}
