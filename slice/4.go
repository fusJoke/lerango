package main

import "fmt"

func main() {
	arr := []int{1,2,3,4}
	fmt.Println(arr)
	fmt.Println(&arr)
	changeArr(arr)
	fmt.Println(arr)
	fmt.Println(&arr)
}

func changeArr(arr []int){
	arr = append(arr , 33)
	fmt.Println(&arr)
}
