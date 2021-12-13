package filter

import (
	"io"
	"text/template"

	"github.com/wanglihui/slicer/fns/core"
)

func NewFilterTemplate() *FilterTemplate {
	return &FilterTemplate{}
}

type FilterTemplate struct {
}

func (tpl *FilterTemplate) GetContent(out io.Writer, data core.GenerateData) error {
	return tpl.filterFnTpl(out, data)
}

func (tpl *FilterTemplate) filterFnTpl(out io.Writer, data core.GenerateData) error {
	t, err := template.New("filterFnTpl").Parse(`
type {{.Type}}FilterFn func(item {{.Type}}, idx int) bool
func (items {{.Type}}Slice) Filter(fn {{.Type}}FilterFn) {{.Type}}Slice {
	rets := make({{.Type}}Slice, 0)
	for idx, item := range items {
		if b := fn(item, idx); b {
			rets = append(rets, item)
		}
	}
	return rets
}
	`)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
