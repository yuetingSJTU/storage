package main

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemType := v.Elem()

	n := elemType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemType)
		return
	}

	fmt.Printf("%s's method set:\n", elemType)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemType.Method(j).Name)
	}
	fmt.Printf("\n")
}
