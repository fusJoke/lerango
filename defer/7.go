package main

import "fmt"

func main() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic(" panic")
}

/*
wangfusheng@MacBook-Pro-2 defer % go run 7.go
defer panic
 */