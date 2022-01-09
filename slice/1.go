package main

import "fmt"

func main() {
	var array = [5]int{1,2,3,4,5}

	slice1 := array[1:2]
	slice2 := array[1:4]
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 := append(slice1, 5)
	fmt.Println(slice3)
	fmt.Println(slice2)

}