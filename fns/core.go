package fns

import (
	"io"
	"text/template"
)

type GenerateData struct {
	Type     string
	PkgName  string
	FileName string
	Line     int
}

type Template interface {
	GetContent(out io.Writer, data GenerateData) error
}

func NewCoreTemplate() *CoreTemplate {
	return &CoreTemplate{}
}

type CoreTemplate struct {
}

func (tpl *CoreTemplate) GetContent(out io.Writer, data GenerateData) error {
	t, _ := template.New("core").Parse(`
// Code generated by slicer DO NOT EDIT.
// The source file is {{.FileName}} line number {{.Line}}
// more about to view https://github.com/wanglihui/slicer
package {{.PkgName}}
type {{.Type}}Slice []{{.Type}}
	`)
	return t.Execute(out, data)
}