package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float32 = 3.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t)
	fmt.Println(v)

	var  A interface{}
	A = 100

	v = reflect.ValueOf(A)
	B := v.Interface()

	if v == B {
		fmt.Printf("they are same")
	}


}
