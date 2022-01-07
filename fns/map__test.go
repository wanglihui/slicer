package fns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/slicer/fns"
)

func TestMap(t *testing.T) {
	t.Run("map should be ok", func(t *testing.T) {
		arr := []fns.Template___{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
		}
		arr2 := fns.Template___Slice(arr).MapReturnTemplate___(func(item fns.Template___, idx int) fns.ReturnTemplate___ {
			return fns.ReturnTemplate___{}
		})
		assert.Equal(t, len(arr2), len(arr))
	})

}
