package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/wanglihui/slicer/fns/core"
	"github.com/wanglihui/slicer/fns/filter"
)

var typeName = flag.String("type", "", "input type name")
var pkgName = flag.String("pkg", "", "input package name")

func main() {
	flag.Parse()
	if typeName == nil {
		panic("need set -type")
	}
	if pkgName == nil {
		panic("need set -pkg")
	}
	tpls := []core.Template{
		core.NewCoreTemplate(),
		filter.NewFilterTemplate(),
	}
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	f, err := os.Create(path.Join(dir, fmt.Sprintf("./%s_slice.gen.go", strings.ToLower(*typeName))))
	if err != nil {
		panic(err)
	}
	data := core.GenerateData{
		Type:    *typeName,
		PkgName: *pkgName,
	}
	for _, tpl := range tpls {
		tpl.GetContent(f, data)
	}
}
