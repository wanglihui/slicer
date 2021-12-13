package fns

import (
	"github.com/wanglihui/slicer/fns/core"
	"github.com/wanglihui/slicer/fns/filter"
)

func NewFnTemplate(core *core.CoreTemplate, filter *filter.FilterTemplate) *FnTemplate {
	return &FnTemplate{
		Core:   core,
		Filter: filter,
	}
}

type FnTemplate struct {
	Core   *core.CoreTemplate
	Filter *filter.FilterTemplate
}

// func (FnTemplate) GetContent(name string) string {

// }
