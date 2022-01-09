package main

import "fmt"

func selectSort(list []int) {
	len := len(list)

	for i :=0; i < len-1; i++ {
		min := list[i]
		minIndex := i
		for j := i + 1; j < len; j++ {
			if list[j] < min {
				if list[j] < min {
					min = list[j]
					minIndex = j
				}
			}
		}

		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	selectSort(list)
	fmt.Println(list)
}