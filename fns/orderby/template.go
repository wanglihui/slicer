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
	return quickSort(items, 0, len(items)-1, fn)
}
func quickSort(arr {{.Type}}Slice, left, right int, fn {{.Type}}CompareFn) {{.Type}}Slice {
	if left < right {
			partitionIndex := partition(arr, left, right, fn)
			quickSort(arr, left, partitionIndex-1, fn)
			quickSort(arr, partitionIndex+1, right, fn)
	}
	return arr
}

func partition(arr {{.Type}}Slice, left, right int, fn {{.Type}}CompareFn) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
			if fn(arr[i], arr[pivot]) < 0 {
					swap(arr, i, index)
					index += 1
			}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr {{.Type}}Slice, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
	`)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
