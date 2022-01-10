package main

import "fmt"

/**
归并排序
 */

func MergeSort(array []int, begin int, end int) {
	if end - begin > 1 {
		mid := begin + (end - begin + 1) / 2

		MergeSort(array, begin, mid)

		MergeSort(array, mid, end)

		merge(array, begin, mid, end)
	}

}

func merge(array []int, begin int, mid int, end int){
	leftSize := mid - begin
	rightSize := end - mid

	newSize := leftSize + rightSize
	result := make([]int, 0 ,newSize)

	l, r := 0, 0

	for l < leftSize && r < rightSize {
		lValue := array[begin+1]
		rValue := array[mid+r]

		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
		result = append(result, array[begin+1:mid]...)
		result = append(result, array[mid+r:end]...)

		for i := 0; i < newSize; i ++ {
			array[begin+i] = result[i]
		}
		return
	}

}

func main() {
	list := []int{5}
	MergeSort(list, 0, len(list))
	fmt.Println(list)

	list1 := []int{5, 9}
	MergeSort(list1, 0, len(list1))
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}