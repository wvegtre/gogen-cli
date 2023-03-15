package dbconverter

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

//func WithEnableJsonTag(enable bool) MySQLConfigOption {
//	return func(config *mySQLConverterConfig) {
//		config.EnableJsonTag = enable
//	}
//}

func TestGenOptionsByConfig(t *testing.T) {
	var config mySQLConverterConfig
	rt := reflect.TypeOf(config)
	rv := reflect.ValueOf(config)
	if rt.Kind() != reflect.Struct {
		t.Error("must be a struct")
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		if rv.Field(i).Type().Kind() == reflect.Struct {
			internalField := rv.Field(i).Type()
			for j := 0; j < internalField.NumField(); j++ {
				//t.Log(internalField.Field(j).Name)
				printOptionsFuncWithName(internalField.Field(j).Name)
			}
		} else {
			//t.Log(rt.Field(i).Name)
			printOptionsFuncWithName(rt.Field(i).Name)
		}
	}
}

var temp = `func With%s(%s string) MySQLConfigOption {
	return func(config *mySQLConverterConfig) {
		config.%s = %s
	}
}
`

func printOptionsFuncWithName(name string) {
	firstUpper := strings.ToUpper(name[:1]) + name[1:]
	firstLower := strings.ToLower(name[:1]) + name[1:]

	content := fmt.Sprintf(temp, firstUpper, firstLower, firstUpper, firstLower)
	fmt.Println(content)
}
