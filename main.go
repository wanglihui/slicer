package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/wanglihui/slicer/fns/core"
	"github.com/wanglihui/slicer/fns/filter"
	"github.com/wanglihui/slicer/fns/find"
	"github.com/wanglihui/slicer/fns/maps"
	"github.com/wanglihui/slicer/fns/orderby"
)

var typeName = flag.String("type", "", "input type name")

func main() {
	flag.Parse()
	if typeName == nil {
		panic("need set -type")
	}
	pkgName := os.Getenv("GOPACKAGE")
	fileName := os.Getenv("GOFILE")
	line := os.Getenv("GOLINE")
	tpls := []core.Template{
		core.NewCoreTemplate(),
		filter.NewFilterTemplate(),
		maps.NewMapTemplate(),
		orderby.NewOrderByTemplate(),
		find.NewFindTemplate(),
	}
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	f, err := os.Create(path.Join(dir, fmt.Sprintf("./%s_slice.gen.go", strings.ToLower(*typeName))))
	if err != nil {
		panic(err)
	}
	linenum, _ := strconv.Atoi(line)
	data := core.GenerateData{
		Type:     *typeName,
		PkgName:  pkgName,
		FileName: fileName,
		Line:     linenum,
	}
	for _, tpl := range tpls {
		tpl.GetContent(f, data)
	}
}
