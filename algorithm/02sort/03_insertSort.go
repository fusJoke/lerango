package main

import "fmt"

func InsertSort(list []int) {
	n := len(list)
	for i := 1; i <= n-1; i++ {
		deal := list[i]
		j := i - 1

		if deal < list[j] {
			for ;j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] =deal
		}

	}
}

func main() {
	list := []int{5}
	InsertSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	InsertSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list2)
	fmt.Println(list2)
}