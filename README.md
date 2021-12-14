# slicer

use go generate to generate slice helper code. like map, find, groupby, orderby, find, first and so on .

### Usage

```golang
//go:generate go run github.com/wanglihui/slicer -type PackageItem -pkg packageitem
```
- type is struct type
- pkg is package name

```bash
go generate ./...
```

now go generate output code like 
you can use Map Find Filter OrderBy and so on...

```golang
// packageitem_slice.gen.go
package packageitem

type PackageItem interface{}

type PackageItemSlice []PackageItem

type PackageItemFilterFn func(item PackageItem, idx int) bool

func (items PackageItemSlice) Filter(fn PackageItemFilterFn) PackageItemSlice {
	rets := make(PackageItemSlice, 0)
	for idx, item := range items {
		if b := fn(item, idx); b {
			rets = append(rets, item)
		}
	}
	return rets
}

func (items PackageItemSlice) OrderBy(fn PackageItemCompareFn) PackageItemSlice {
	return _PackageItemQuickSort(items, 0, len(items)-1, fn)
}

type PackageItemFindFn func(item PackageItem) bool

func (items PackageItemSlice) Find(fn PackageItemFindFn) *PackageItem {
	for _, item := range items {
		if b := fn(item); b {
			return &item
		}
	}
	return nil
}
// and so on

```

```golang
//main.go

// Map
PackageItemSlice.Map(func(item PackageItem) interface{} {
    return item.Name
})
// Find one
item := PackageItemSlice.Find(func(item PackageItem) *PackageItem {
    return &item.Age > 10
})

// Filter
items := PackageItemSlice.Filter(func(item PackageItem) bool {
    return item.Age < 10
})
```