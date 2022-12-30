package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

var csharpTypeNames map[string]string
var golangTypeNames map[string]string

type exporter struct {
	luaTableTemplate []byte
	jsonTemplate     []byte
	csharpTemplate   []byte
	golangTemplate   []byte
}

func (e *exporter) Init() {
	e.luaTableTemplate, _ = ioutil.ReadFile("templates/lua.tmpl")
	e.jsonTemplate, _ = ioutil.ReadFile("templates/json.tmpl")
	e.csharpTemplate, _ = ioutil.ReadFile("templates/csharp.tmpl")
	e.golangTemplate, _ = ioutil.ReadFile("templates/golang.tmpl")

	initCsharpTypeNames()
	initGolangTypeNames()

}

func (e *exporter) ExportLua(folder string, xl *xlsx) {
	tmpl, err := template.New("luaExport").Funcs(genericFuncMap()).Parse(string(e.luaTableTemplate))
	if err != nil {
		panic(err)
	}
	newFileName := path.Join(folder, xl.SubPath)
	if _, err := os.Stat(newFileName); os.IsNotExist(err) {
		os.MkdirAll(newFileName, os.ModePerm) // mkdir 创建目录，mkdirAll 可创建多层级目录
	}
	newFileName = path.Join(newFileName, xl.FileName+".lua")
	newFile, err := os.Create(newFileName)
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xl)
	if err != nil {
		panic(err)
	}
}

func (e *exporter) ExportJSON(folder string, xl *xlsx) {
	tmpl, err := template.New("jsonExport").Funcs(genericFuncMap()).Parse(string(e.jsonTemplate))
	if err != nil {
		panic(err)
	}
	newFileName := path.Join(folder, xl.SubPath)
	if _, err := os.Stat(newFileName); os.IsNotExist(err) {
		os.MkdirAll(newFileName, os.ModePerm) // mkdir 创建目录，mkdirAll 可创建多层级目录
	}
	newFileName = path.Join(newFileName, xl.FileName+".json")
	newFile, err := os.Create(newFileName)
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xl)
	if err != nil {
		panic(err)
	}
}

func (e *exporter) ExportCSharp(folder string, xl *xlsx) {
	tmpl, err := template.New("csharpExport").Funcs(genericFuncMap()).Parse(string(e.csharpTemplate))
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xl.Name+".cs"))
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(newFile, xl)
	if err != nil {
		panic(err)
	}
}

func (e *exporter) ExportGolang(folder string, xl *xlsx) {
	tmpl, err := template.New("golangExport").Funcs(genericFuncMap()).Parse(string(e.golangTemplate))
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(path.Join(folder, xl.FileName+".go"))
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, xl); err != nil {
		panic(err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	newFile.Write(p)
}

func genericFuncMap() map[string]interface{} {
	var genericMap = map[string]interface{}{
		"GetCSharpTypeName": parseCsharpType,
		"GetGolangTypeName": parseGolangType,
		"indent":            indent,
		"csharpInherit":     csharpInherit,
		"golangInherit":     golangInherit,
	}

	return genericMap
}

func parseCsharpType(longType string) string {
	if csharpTypeNames[longType] != "" {
		return csharpTypeNames[longType]
	}
	first := strings.Index(longType, "<")
	last := strings.LastIndex(longType, ">")
	if first != -1 && last != -1 {
		thistype := longType[:first]
		subType := longType[first+1 : last]
		return csharpTypeNames[thistype] + parseCsharpType(subType) + ">"
	}
	return ""
}

func parseGolangType(longType string) string {
	if golangTypeNames[longType] != "" {
		return golangTypeNames[longType]
	}
	first := strings.Index(longType, "<")
	last := strings.LastIndex(longType, ">")
	if first != -1 && last != -1 {
		thistype := longType[:first]
		subType := longType[first+1 : last]
		return golangTypeNames[thistype] + parseGolangType(subType)
	}
	return ""
}

func initCsharpTypeNames() {
	csharpTypeNames = make(map[string]string)
	csharpTypeNames["int"] = "int"
	csharpTypeNames["float"] = "float"
	csharpTypeNames["string"] = "string"
	csharpTypeNames["bool"] = "bool"
	csharpTypeNames["dict"] = "Dictionary<string, "
	csharpTypeNames["list"] = "List<"
}

func initGolangTypeNames() {
	golangTypeNames = make(map[string]string)
	golangTypeNames["int"] = "int32"
	golangTypeNames["float"] = "float32"
	golangTypeNames["string"] = "string"
	golangTypeNames["bool"] = "bool"
	golangTypeNames["dict"] = "map[string]"
	golangTypeNames["list"] = "[]"
}
