package main

import "fmt"

// QuickSort 普通快速排序
func QuickSort(array []int, begin, end int) {
	if begin < end {

		loc := partition(array, begin, end)

		QuickSort(array, begin, loc-1)

		QuickSort(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	i := begin + 1
	j := end

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	if array[i] >= array[begin] {
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	list := []int{5}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)

	list1 := []int{5, 9}
	QuickSort(list1, 0, len(list1)-1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	QuickSort(list2, 0, len(list2)-1)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}