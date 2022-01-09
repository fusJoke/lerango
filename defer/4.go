package main

import "fmt"

func main() {
	x := defer3func()
	fmt.Println(x)		// x == 10
}

func defer3func() (t int){
	defer func() {
		t *= 10
	}()
	return 1
}
