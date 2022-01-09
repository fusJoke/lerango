package main

import "fmt"

func main() {
	var nums = []int{1,2,3,4,5,6,8,9,12,13,16,21,24}
	num := search(nums,12)
	fmt.Println(num)
}

func search(nums []int, target int) int {
	low, high := 0, len(nums) -1
	for low <= high {
		mid := (high -low) /2
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			low = mid - 1
		} else {
			high = mid + 1
		}
	}
	return -1
}




