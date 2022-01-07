package fns

type Template___FindFn func(item Template___) bool

func (items Template___Slice) Find(fn Template___FindFn) *Template___ {
	for _, item := range items {
		if b := fn(item); b {
			return &item
		}
	}
	return nil
}
