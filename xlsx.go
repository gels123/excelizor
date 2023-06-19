package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type testgels struct {
	num int32
}

type xlsx struct {
	Name        string
	FileName    string
	Template    *xField
	OwnTemplate *xField
	Data        []*xField
	keymap      map[int]*xField
	ParentClass string
	SubPath     string
}

func (x *xlsx) Init(fileName string, name string) {
	x.Name = name
	x.FileName = fileName
	x.Data = make([]*xField, 0)
	x.keymap = make(map[int]*xField)
}

func (x *xlsx) Parse(rows [][]string) {
	x.Template = new(xField)
	x.OwnTemplate = x.Template.Copy()
	if ok, _ := x.Template.Init(x.Name, "struct", ""); ok {
		x.Template.ParseSubFieldsDefs(rows[1], rows[2], rows[3])

		for i := 4; i < len(rows); i++ {
			field := x.Template.Copy()

			// comment row
			if strings.HasPrefix(rows[i][0], "//") || rows[i][0] == "" {
				continue
			}
			id, _ := strconv.Atoi(rows[i][0])
			if _, ok2 := x.keymap[id]; !ok2 {
				field.ParseDatas(id, rows[i])
				field.SetLevel(4)
				x.Data = append(x.Data, field)
				x.keymap[id] = field
			} else {
				log.Fatalln("Parse", x.Name, "failed, Id", id, "is duplicated")
			}
		}
		i := 0
		for i < len(x.Template.Fields) {
			v := x.Template.Fields[i]
			if strings.HasPrefix(v.Type, "//") || (v.Tag != "" && v.Tag != params.tag) {
				x.Template.Fields = append(x.Template.Fields[:i], x.Template.Fields[i+1:]...)
			} else {
				i++
			}
		}

		if x.ParentClass != "" {
			x.OwnTemplate.Fields = x.OwnTemplate.Fields[:0]
			parentClassName := x.ParentClass + ".xlsx"
			if val, ok := loadedFiles[parentClassName]; ok {
				parseExcel(parentClassName, val)
			} else {
				log.Fatalln("Cant find parent class excel name = " + parentClassName)
			}
			for index := 0; index < len(x.Template.Fields); index++ {
				field := x.Template.Fields[index]
				if !loadedFiles[parentClassName].xl.CheckExistField(field) {
					x.OwnTemplate.Fields = append(x.OwnTemplate.Fields, field.Copy())
				}
			}

		} else {
			x.OwnTemplate = x.Template.Copy()
		}
	} else {
		log.Fatalln("Parse", x.Name, "head field")
	}
}

func (x *xlsx) CheckExistField(field *xField) bool {
	for index := 0; index < len(x.Template.Fields); index++ {
		if x.Template.Fields[index].Name == field.Name && x.Template.Fields[index].LongType == field.LongType {
			return true
		}
	}
	return false
}

func (x *xlsx) Print() {
	for k, v := range x.Data {
		fmt.Print(k, " ")
		v.Print()
	}
}
