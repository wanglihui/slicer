package fns

func (that Template___Slice) Chunk(chunkSize int) [][]Template___ {
	if len(that) == 0 {
		return nil
	}
	divided := make([][]Template___, (len(that)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(that) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = that[prev:next]
		prev = next
		i++
	}
	divided[i] = that[prev:]
	return divided
}
