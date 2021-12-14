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
type {{.Type}}MapFn func(item {{.Type}}, idx int) interface{}
type {{.Type}}Map{{.Type}}Fn func(item {{.Type}}, idx int) {{.Type}}
type {{.Type}}MapStringFn func(item {{.Type}}, idx int) string
type {{.Type}}MapIntFn func(item {{.Type}}, idx int) int
type {{.Type}}MapInt64Fn func(item {{.Type}}, idx int) int64
type {{.Type}}MapFloat64Fn func(item {{.Type}}, idx int) float64
type {{.Type}}MapReturnedInterface []interface{}

func (items {{.Type}}Slice) Map(fn {{.Type}}MapFn) {{.Type}}MapReturnedInterface {
	arr := make([]interface{}, 0)
	for idx, item := range items {
		arr = append(arr, fn(item, idx))
	}
	return arr
}

func (items {{.Type}}Slice) MapString(fn {{.Type}}MapStringFn) []string {
	rets := make([]string, 0)
	for idx, item := range items {
		s := fn(item, idx)
		rets = append(rets, s)
	}
	return rets
}

func (items {{.Type}}Slice) MapInt(fn {{.Type}}MapIntFn) []int {
	rets := make([]int, 0)
	for idx, item := range items {
		s := fn(item, idx)
		rets = append(rets, s)
	}
	return rets
}

func (items {{.Type}}Slice) MapInt64(fn {{.Type}}MapInt64Fn) []int64 {
	rets := make([]int64, 0)
	for idx, item := range items {
		s := fn(item, idx)
		rets = append(rets, s)
	}
	return rets
}

func (items {{.Type}}Slice) MapFloat64(fn {{.Type}}MapFloat64Fn) []float64 {
	rets := make([]float64, 0)
	for idx, item := range items {
		s := fn(item, idx)
		rets = append(rets, s)
	}
	return rets
}

func (items {{.Type}}Slice) Map{{.Type}}(fn {{.Type}}Map{{.Type}}Fn) []{{.Type}} {
	rets := make([]{{.Type}}, 0)
	for idx, item := range items {
		s := fn(item, idx)
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
