package fns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/slicer/fns"
)

func TestFilter(t *testing.T) {
	t.Run("filter return false", func(t *testing.T) {
		arr := []fns.Template___{
			{},
			{},
			{},
			{},
		}
		arr2 := fns.Template___Slice(arr).Filter(func(item fns.Template___, idx int) bool {
			return false
		})
		assert.Equal(t, len(arr2), 0)
	})

	t.Run("filter return true", func(t *testing.T) {
		arr := []fns.Template___{
			{},
			{},
			{},
			{},
		}
		arr2 := fns.Template___Slice(arr).Filter(func(item fns.Template___, idx int) bool {
			return true
		})
		assert.Equal(t, len(arr2), len(arr))
	})

}
