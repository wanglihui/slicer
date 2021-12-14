package find

import (
	"html/template"
	"io"

	"github.com/wanglihui/slicer/fns/core"
)

func NewFindTemplate() *FindTemplate {
	return &FindTemplate{}
}

type FindTemplate struct{}

func (tpl *FindTemplate) GetContent(out io.Writer, data core.GenerateData) error {
	t, err := template.New("find").Parse(
		`
type {{.Type}}FindFn func(item {{.Type}}) bool
func (items {{.Type}}Slice) Find(fn {{.Type}}FindFn) *{{.Type}} {
	for _, item := range items {
		if b := fn(item); b {
			return &item
		}
	}
	return nil
}
		`)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
