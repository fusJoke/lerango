package main

import "fmt"

/**
	初始化的方式有两种
	1. 变量声明
	2. 使用内置函数make()
 */
var ch chan int
func main() {
	ch = make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan string, 5)
	ch1 <- 1
	d := <-ch1
	fmt.Println(d)

	ch2 <- "sss"
	s2 := <- ch2

	fmt.Println(s2)
}