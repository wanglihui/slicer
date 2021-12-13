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
type mapFn func(pkg {{.Type}}, idx int) interface{}
type mapStringFn func(pkg {{.Type}}, idx int) string
type mapIntFn func(pkg {{.Type}}, idx int) int
type mapInt64Fn func(pkg {{.Type}}, idx int) int64
type mapFloat64Fn func(pkg {{.Type}}, idx int) float64
type MapReturnedInterface []interface{}

func (packages {{.Type}}Slice) Map(fn mapFn) MapReturnedInterface {
	arr := make([]interface{}, 0)
	for idx, pkg := range packages {
		arr = append(arr, fn(pkg, idx))
	}
	return arr
}

func (packages {{.Type}}Slice) MapString(fn mapStringFn) []string {
	rets := make([]string, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapInt(fn mapIntFn) []int {
	rets := make([]int, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapInt64(fn mapInt64Fn) []int64 {
	rets := make([]int64, 0)
	for idx, pkg := range packages {
		s := fn(pkg, idx)
		rets = append(rets, s)
	}
	return rets
}

func (packages {{.Type}}Slice) MapFloat64(fn mapFloat64Fn) []float64 {
	rets := make([]float64, 0)
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
