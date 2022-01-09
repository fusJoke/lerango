package main

import "fmt"

func main() {
	var  i = 1
	var str string
	str = "sdsdsd"
	doSomething(i)
	doSomething(str)
}

func doSomething(i interface{}) {
	switch v:= i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("i don't know")
	}
}