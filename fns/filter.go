package fns

type Template___FilterFn func(item Template___, idx int) bool

func (items Template___Slice) Filter(fn Template___FilterFn) Template___Slice {
	rets := make(Template___Slice, 0)
	for idx, item := range items {
		if b := fn(item, idx); b {
			rets = append(rets, item)
		}
	}
	return rets
}
