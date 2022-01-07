package fns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/slicer/fns"
)

func TestFind(t *testing.T) {
	t.Run("find has", func(t *testing.T) {
		arr := []fns.Template___{
			{},
			{},
			{},
			{},
		}
		it := fns.Template___Slice(arr).Find(func(item fns.Template___) bool {
			return false
		})
		assert.Empty(t, it)
	})

	t.Run("find not has", func(t *testing.T) {
		arr := []fns.Template___{
			{},
			{},
			{},
			{},
		}
		it := fns.Template___Slice(arr).Find(func(item fns.Template___) bool {
			return true
		})
		assert.NotNil(t, it)
	})

}
