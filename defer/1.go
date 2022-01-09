package main

import "fmt"

/**
	多个defer的执行顺序
 */

func main() {
	defer func1()
	defer func2()
	defer func3()
}

func func1(){
	fmt.Println("func1")
}

func func2(){
	fmt.Println("func2")
}

func func3(){
	fmt.Println("func3")
}

/**
wangfusheng@MacBook-Pro-2 defer % go run 1.go
func3
func2
func1
 */
