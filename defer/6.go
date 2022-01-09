package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		fmt.Println("defer2 before panic")
	}()

	panic("panic do")

	defer func() {
		fmt.Println("defer2 after panic")
	}()	//这部分无法被执行到
}
/**
defer 中用recover接收
 */


