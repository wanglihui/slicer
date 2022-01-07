package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
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
	tpls := []string{
		"chunk",
		"map",
		"filter",
		"find",
		"orderby",
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
	returnTypes := map[string][]string{
		"map": []string{"int", "string", "int64"},
	}
	data := &GenerateData{
		Type:        *typeName,
		PkgName:     pkgName,
		FileName:    fileName,
		Line:        linenum,
		ReturnTypes: returnTypes,
	}
	generateHeadBrand(f, data)
	for _, tpl := range tpls {
		generateFn(f, tpl, data)
	}
}

func generateHeadBrand(out io.Writer, data *GenerateData) error {
	t, _ := template.New("core").Parse(`
// Code generated by slicer DO NOT EDIT.
// The source file is {{.FileName}} line number {{.Line}}
// more about to view https://github.com/wanglihui/slicer
package {{.PkgName}}
type {{.Type}}Slice []{{.Type}}
	`)
	return t.Execute(out, data)
}

type GenerateData struct {
	Type        string
	PkgName     string
	FileName    string
	Line        int
	ReturnTypes map[string][]string
}

func generateFn(out io.Writer, fn string, data *GenerateData) error {
	_, filename, _, _ := runtime.Caller(1)
	filesrc := path.Join(path.Dir(filename), "fns/"+fn+".go")
	content, err := ioutil.ReadFile(filesrc)
	if err != nil {
		return err
	}
	s := string(content)
	if data.ReturnTypes != nil && len(data.ReturnTypes[fn]) > 0 {
		s = handleReturnTypes(s, data.ReturnTypes[fn], data)
	}
	s = strings.Replace(s, "Template___", data.Type, -1)
	s = strings.Replace(s, "package fns", "", -1)
	if _, err := out.Write([]byte(s)); err != nil {
		return err
	}
	return nil
}

func handleReturnTypes(content string, returnTypes []string, data *GenerateData) string {
	s := ""
	for _, returnType := range returnTypes {
		s += strings.Replace(content, "ReturnTemplate___", returnType, -1)
	}
	return s
}
