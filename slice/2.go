package main

import "fmt"

func main() {
	data := [...]int{1,2,3,4,5,6,7,8}

	s := data[0:4]
	s[0] = 100
	s[1] = 200
	fmt.Println(data)
	fmt.Println(s)
}