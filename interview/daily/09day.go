package main

import "fmt"

func hello(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}	//切片是引用
	hello(i...)
	fmt.Println(i[0])
}