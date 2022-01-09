package main

import "fmt"

func returnButDefer() (t int) {  //t初始化0， 并且作用域为该函数全域
	x := 1
	defer func() {
		t = t * 10
	}()

	return x
}

func main() {
	fmt.Println(returnButDefer())
}