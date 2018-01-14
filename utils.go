package main

import (
	"errors"
	"strconv"
	"strings"
)

func trimData(data string) string {
	before := data
	for {
		data = strings.TrimPrefix(data, "{")
		data = strings.TrimSuffix(data, "}")
		if before == data {
			return data
		}
		before = data
	}

}

func splitName(data string) []string {
	subDatas := make([]string, 0)
	equal := strings.Index(data, "=")

	var n string
	var d string
	if equal == -1 {
		d = data
	} else {
		n = data[:equal]
		d = data[equal+1:]
	}
	subDatas = append(subDatas, n)
	subDatas = append(subDatas, d)
	return subDatas
}

func splitSubData(layer int, data string) []string {
	sept := ""
	for i := 1; i < layer; i++ {
		sept += "}"
	}
	sept += "|"

	subDatas := make([]string, 0)

	for {
		pos := strings.Index(data, sept)
		if pos == -1 {
			subDatas = append(subDatas, data)
			break
		} else {
			subData := data[0 : pos+layer-1]
			data = data[pos+layer:]
			subDatas = append(subDatas, subData)
		}
	}
	return subDatas
}

func handleData(dataType string, data string) (string, error) {
	var result string
	var retErr error
	switch dataType {
	case "int":
		ret, err := strconv.Atoi(data)
		result = strconv.Itoa(ret)
		retErr = err
	case "float":
		ret, err := strconv.ParseFloat(data, 32)
		result = strconv.FormatFloat(ret, 'f', 3, 32)
		retErr = err
	case "bool":
		ret, err := strconv.ParseBool(data)
		result = strconv.FormatBool(ret)
		retErr = err
	case "string":
		result = data
		retErr = nil
	default:
		retErr = errors.New("DataType " + dataType + " is invalid for data " + data)
	}
	return result, retErr
}

func name2lower2Camel(name string) (string, string) {
	dotIndex := strings.LastIndex(name, ".")
	lower := name[:dotIndex]

	initial := strings.ToUpper(lower[0:1])
	other := lower[1:]
	for strings.Index(other, "_") != -1 {
		index := strings.Index(other, "_")
		replace := strings.ToUpper(other[index+1 : index+2])
		s := []string{other[:index], replace, other[index+2:]}
		other = strings.Join(s, "")
	}
	return lower, initial + other
}

func upperInitialChar(str string) string {
	initial := strings.ToUpper(str[0:1])
	other := str[1:]
	return initial + other
}
