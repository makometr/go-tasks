package task17

import (
	"fmt"
	"reflect"
)

// Написать программу, которая в рантайме способна определить тип переменной — int, string, bool, channel из переменной типа interface{}.

// Task17 prints type
func Task17(x interface{}) {
	switch reflect.ValueOf(x).Kind() {
	case reflect.Bool:
		fmt.Printf("bool")
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Printf("int")
	case reflect.String:
		fmt.Printf("string")
	case reflect.Chan:
		fmt.Printf("chan")
	default:
		fmt.Println("undefined")
	}

}
