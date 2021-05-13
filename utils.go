package utils

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

var cyan func(a ...interface{}) string
var succes func(a ...interface{}) string
var red func(a ...interface{}) string

func init() {
	cyan = color.New(color.FgCyan).SprintFunc()
	succes = color.New(color.FgGreen).SprintFunc()
	red = color.New(color.FgRed).SprintFunc()
}

func PreetyPrint(modifier string, inputs ...interface{}) {
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		in := reflect.ValueOf(input)

		if in.Kind() == reflect.Struct {
			numfields := in.NumField()
			fmt.Printf("%v%v: %v", modifier, in.Type().Name(), red("{"))
			for nm := 0; nm < numfields; nm++ {
				field := in.Field(nm)
				if field.Kind() == reflect.Struct {
					// fmt.Print("\n")
					PreetyPrint(fmt.Sprintf("%v", modifier), field.Interface())
				} else {
					fmt.Printf("%v:%v, ", in.Type().Field(nm).Name, cyan(field.Interface()))
				}
			}
			fmt.Printf(red("}"))
		} else {
			fmt.Printf("%vT: %v; V: %v\n", modifier, succes(in.Type()), cyan(i))
		}
	}
}

func SPrintf(s string, inputs ...interface{}) {

}

func SPrint(input interface{}) string {
	in := reflect.ValueOf(input)
	ret := ""
	if in.Kind() == reflect.Struct {
		numfields := in.NumField()
		ret += fmt.Sprintf("%v: %v", in.Type().Name(), red("{"))
		for nm := 0; nm < numfields; nm++ {
			field := in.Field(nm)
			if field.Kind() == reflect.Struct {
				// fmt.Print("\n")
				ret += SPrint(field.Interface())
			} else {
				fmt.Printf("%v:%v, ", in.Type().Field(nm).Name, cyan(field.Interface()))
			}
		}
		fmt.Printf(red("}"))
	} else {
		fmt.Printf("%v:\t %v\n", succes(in.Type()), cyan(input))
	}
	return ret
}
