package utils

import (
	"fmt"
	"math/rand"
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

	colors = make(map[int]func(a ...interface{}) string, 15)
	colors[0] = color.New(color.FgYellow).Add(color.Italic).SprintFunc()
	colors[1] = color.New(color.FgBlue).Add(color.Italic).SprintFunc()
	colors[2] = color.New(color.FgMagenta).Add(color.Italic).SprintFunc()
	colors[3] = color.New(color.FgHiGreen).Add(color.Italic).SprintFunc()

	colors[4] = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	colors[5] = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	colors[6] = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	colors[7] = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()

	colors[8] = color.New(color.FgYellow).Add(color.Bold, color.Italic).SprintFunc()
	colors[9] = color.New(color.FgBlue).Add(color.Bold, color.Italic).SprintFunc()
	colors[10] = color.New(color.FgMagenta).Add(color.Bold, color.Italic).SprintFunc()
	colors[11] = color.New(color.FgHiGreen).Add(color.Bold, color.Italic).SprintFunc()

	colors[12] = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	colors[13] = color.New(color.FgHiRed).Add(color.Bold, color.Italic).SprintFunc()
	colors[14] = color.New(color.FgRed).Add(color.Italic).SprintFunc()
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
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		fmt.Printf(s, SPrint(input))
	}

}

var colors map[int]func(a ...interface{}) string

func SPrint(input interface{}) string {
	in := reflect.ValueOf(input)
	ret := ""
	if in.Kind() == reflect.Struct {
		numfields := in.NumField()
		rcolor := colors[rand.Intn(14)]
		ret += fmt.Sprintf("%v: %v", rcolor(in.Type().Name()), rcolor("{"))

		for nm := 0; nm < numfields; nm++ {
			field := in.Field(nm)

			switch field.Kind() {
			case reflect.Struct:
				ret += SPrint(field.Interface()) + "\n"
			case reflect.Array:
			case reflect.Slice:
				ret += rcolor(in.Type().Field(nm).Name)
				ret += ": " + rcolor("[")
				for i := 0; i < field.Len(); i++ {
					item := field.Index(i)
					ret += SPrint(item.Interface())
					ret += ","
				}
				ret += rcolor("]")
			default:
				ret += fmt.Sprintf("%v:%v, ", in.Type().Field(nm).Name, cyan(field.Interface()))
			}
		}
		ret += rcolor("}")
	} else {
		fmt.Printf("%v:\t %v\n", succes(in.Type()), cyan(input))
	}
	return ret
}
