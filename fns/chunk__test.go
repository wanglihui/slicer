package fns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanglihui/slicer/fns"
)

func TestChunk(t *testing.T) {
	arr := []fns.Template___{
		{},
		{},
		{},
		{},
	}
	arr2 := fns.Template___Slice(arr).Chunk(3)
	assert.Equal(t, len(arr2), 2)
	assert.Equal(t, len(arr2[0]), 3)
	assert.Equal(t, len(arr2[1]), 1)
}
