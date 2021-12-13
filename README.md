# slicer

### Usage

```golang
//go:generate slicer -type PackageItem -pkg main
```

```golang
// packageitem_slice.gen.go
package packageitem

type PackageItemSlice []PackageItem

type MapFn func(pkg PackageItem, idx int) interface{}
type MapPackageItemFn func(pkg PackageItem, idx int) PackageItem
type MapStringFn func(pkg PackageItem, idx int) string
type MapIntFn func(pkg PackageItem, idx int) int
type MapInt64Fn func(pkg PackageItem, idx int) int64
type MapFloat64Fn func(pkg PackageItem, idx int) float64

type FilterFn func(pkg PackageItem, idx int) PackageItem
// type ReduceFn func()
// type GroupByFn func()
// type FirstFn func()
// type Take func()

func (packages PackageItemSlice) Map(fn MapFn) []interfafce{} {
    rets := make([]interface{}, 0)
    for pkg, idx := range packages {
        rets = append(rets, fn(pkg,idx))
    }
    return rets
}
// and so on
```