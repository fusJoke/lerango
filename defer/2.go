package main

import "fmt"

func main() {
	returnAndDefer()
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func returnFunc() int {
	fmt.Println("return do")
	return 0
}

func deferFunc(){
	fmt.Println("deferfunc do")

}

/*
wangfusheng@MacBook-Pro-2 defer % go run 2.go
return do
deferfunc do
 */