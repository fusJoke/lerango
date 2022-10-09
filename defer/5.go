package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer1 before panic")
	}()

	defer func() {
		fmt.Println("defer2 before panic")
	}()

	panic("panic do")

	defer func() {
		fmt.Println("defer2 after panic")
	}() //这部分无法被执行到
}

/**
defer2 before panic
defer1 before panic
panic: panic do

goroutine 1 [running]:
main.main()
	/Users/wangfusheng/go-test/learngo/defer/5isPalindrome.go:14 +0x68
exit status 2
*/
