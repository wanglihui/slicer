package core

import (
	"io"
	"text/template"
)

type GenerateData struct {
	Type    string
	PkgName string
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
// generate by slicer. DON'T EDIT ANY MORE!!!
package {{.PkgName}}
type {{.Type}}Slice []{{.Type}}
	`)
	return t.Execute(out, data)
}
