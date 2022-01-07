package fns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/slicer/fns"
)

func TestOrderBy(t *testing.T) {
	t.Run("orderby return asc", func(t *testing.T) {
		arr := []fns.Template___{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
		}
		arr2 := fns.Template___Slice(arr).OrderBy(func(left, right fns.Template___) int {
			return -(left.ID - right.ID)
		})
		assert.Greater(t, arr2[0].ID, arr2[1].ID)
	})

	t.Run("orderby return esc", func(t *testing.T) {
		arr := []fns.Template___{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
		}
		arr2 := fns.Template___Slice(arr).OrderBy(func(left, right fns.Template___) int {
			return (left.ID - right.ID)
		})
		assert.Greater(t, arr2[1].ID, arr2[0].ID)
	})

}
