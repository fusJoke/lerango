package main

import "fmt"

func main(){
	fmt.Println("vim-go")
	Go(func () {
		fmt.Println("hello")
		panic("一路向西")
	})
}

func Go(x func()) {
	go func(){
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		x()
	}()
}