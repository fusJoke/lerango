package main

import "fmt"

func main() {
	defer2Func(2) // t = 2
}

func defer2Func(i int) (t int) {
	fmt.Println("t = ",t)
	return 2
}
