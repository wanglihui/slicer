package orderby

import (
	"io"
	"text/template"

	"github.com/wanglihui/slicer/fns/core"
)

func NewOrderByTemplate() *OrderByTemplate {
	return &OrderByTemplate{}
}

type OrderByTemplate struct{}

func (tpl *OrderByTemplate) GetContent(out io.Writer, data core.GenerateData) error {
	t, err := template.New("orderby").Parse(`
type {{.Type}}CompareFn	func(left {{.Type}}, right {{.Type}}) int
func (items {{.Type}}Slice) OrderBy(fn {{.Type}}CompareFn) {{.Type}}Slice {
	return _{{.Type}}QuickSort(items, 0, len(items)-1, fn)
}
func _{{.Type}}QuickSort(arr {{.Type}}Slice, left, right int, fn {{.Type}}CompareFn) {{.Type}}Slice {
	if left < right {
			_{{.Type}}PartitionIndex := _{{.Type}}Partition(arr, left, right, fn)
			_{{.Type}}QuickSort(arr, left, _{{.Type}}PartitionIndex-1, fn)
			_{{.Type}}QuickSort(arr, _{{.Type}}PartitionIndex+1, right, fn)
	}
	return arr
}

func _{{.Type}}Partition(arr {{.Type}}Slice, left, right int, fn {{.Type}}CompareFn) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
			if fn(arr[i], arr[pivot]) < 0 {
					_{{.Type}}Swap(arr, i, index)
					index += 1
			}
	}
	_{{.Type}}Swap(arr, pivot, index-1)
	return index - 1
}

func _{{.Type}}Swap(arr {{.Type}}Slice, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
	`)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
