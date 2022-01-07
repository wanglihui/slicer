package fns

type Template___MapReturnTemplate___Fn func(item Template___, idx int) ReturnTemplate___

func (items Template___Slice) MapReturnTemplate___(fn Template___MapReturnTemplate___Fn) []ReturnTemplate___ {
	arr := make([]ReturnTemplate___, 0)
	for idx, item := range items {
		arr = append(arr, fn(item, idx))
	}
	return arr
}
