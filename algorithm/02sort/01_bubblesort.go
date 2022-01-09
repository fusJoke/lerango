package main

import "fmt"

func BubbleSort(list []int) {
	n := len(list)
	didSwap := false

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap{
			return
		}
	}
}

func main(){
	list := []int{3,51,32,423,2,34,56,78,123,534,2}
	BubbleSort(list)
	fmt.Println(list)
}