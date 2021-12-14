# slicer
golang 切片操作函数集合，提供了Map, Find, Filter, Groupby, OrderBy 等等。

### 使用方式

使用golang generate 方式，对切片对象生成辅助函数.
```golang
//go:generate go run github.com/wanglihui/slicer -type PackageItem -pkg packageitem
```
- type 要追加辅助函数的类型。如PackageItem,那么将为 []PackageItem 生成辅助函数，同时生成别名 PackageItemSlice
- pkg 生成的文件名称默认是 ./lowercase(TypeName)_slice.gen.go ，包名称为 pkg 参数. 如 packageitem_slice.gen.go

```bash
//执行generate命令
go generate ./...
```

```golang
// packageitem_slice.gen.go 自动生成的内容大致如下
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
```

```golang
//在其他模块中使用生成的函数和类型

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

### 注意
- 可以将 *.gen.go 加入到 .gitignore中。在CI中自动生成即可.
- 