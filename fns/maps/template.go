package maps

import (
	"html/template"
	"io"

	"github.com/wanglihui/slicer/fns/core"
)

func NewMapTemplate() *MapTemplate {
	return &MapTemplate{}
}

type MapTemplate struct{}

func (tpl *MapTemplate) GetContent(out io.Writer, data core.GenerateData) error {
	t, err := template.New("").Parse(`
type {{.Type}}MapFn func(pkg {{.Type}}, idx int) interface{}
type {{.Type}}Map{{.Type}}Fn func(pkg {{.Type}}, idx int) {{.Type}}Slice
type {{.Type}}MapStringFn func(pkg {{.Type}}, idx int) string
type {{.Type}}MapIntFn func(pkg {{.Type}}, idx int) int
type {{.Type}}MapInt64Fn func(pkg {{.Type}}, idx int) int64
type {{.Type}}MapFloat64Fn func(pkg {{.Type}}, idx int) float64
type {{.Type}}MapReturnedInterface []interface{}

func (packages {{.Type}}Slice) Map(fn {{.Type}}MapFn) {{.Type}}MapReturnedInterface {
	arr := make([]interface{}, 0)
	for idx, pkg := range packages {
		arr = append(arr, fn(pkg, idx))
	}
	return arr
}

func (packages {{.Type}}Slice) MapString(fn {{.Type}}MapStringFn) []string {
	rets := make([]string, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapInt(fn {{.Type}}MapIntFn) []int {
	rets := make([]int, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapInt64(fn {{.Type}}MapInt64Fn) []int64 {
	rets := make([]int64, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapFloat64(fn {{.Type}}MapFloat64Fn) []float64 {
	rets := make([]float64, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) Map{{.Type}}(fn {{.Type}}Map{{.Type}}Fn) []{{.Type}} {
	rets := make([]{{.Type}}, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}
	`)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
