package fns

type Template___CompareFn func(left Template___, right Template___) int

func (items Template___Slice) OrderBy(fn Template___CompareFn) Template___Slice {
	return _Template___QuickSort(items, 0, len(items)-1, fn)
}
func _Template___QuickSort(arr Template___Slice, left, right int, fn Template___CompareFn) Template___Slice {
	if left < right {
		_Template___PartitionIndex := _Template___Partition(arr, left, right, fn)
		_Template___QuickSort(arr, left, _Template___PartitionIndex-1, fn)
		_Template___QuickSort(arr, _Template___PartitionIndex+1, right, fn)
	}
	return arr
}

func _Template___Partition(arr Template___Slice, left, right int, fn Template___CompareFn) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if fn(arr[i], arr[pivot]) < 0 {
			_Template___Swap(arr, i, index)
			index += 1
		}
	}
	_Template___Swap(arr, pivot, index-1)
	return index - 1
}

func _Template___Swap(arr Template___Slice, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
