package main

import "fmt"

func main() {
	arr := make([]int,10,20)
	arr1 := append(arr,5,6,7)
	fmt.Println(arr)
	_ = append(arr1,5)
	fmt.Println(arr1)
}
