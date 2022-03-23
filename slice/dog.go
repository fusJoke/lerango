package main

import "fmt"

type Dog struct {
	name string
}

var dog Dog
func main() {
	dog1 := dog
	dog2 := new(Dog)

	fmt.Print(dog1)
	fmt.Print(dog2)
}
