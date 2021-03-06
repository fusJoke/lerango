package main

import "fmt"

func ShellSort(list []int)  {
	n := len(list)

	for step := n/2; step >= 1; step /=2 {
		for i := step; i < n; i += step {
			for j := i - step; j >= 0 ; j -= step {
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

func main() {
	list := []int{5}
	ShellSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	ShellSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	ShellSort(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	ShellSort(list3)
	fmt.Println(list3)
}