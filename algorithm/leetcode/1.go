package main

import "fmt"

func main() {
	var nums  = []int{1,2,1,45,1,12,4,3,2}
	var k int = 4
	fmt.Println(numberOfSubArrays(nums,k))
}

func numberOfSubArrays(nums []int, k int) int {
	var hight int = len(nums) - 1
	var count int = 0

	for i := 0 ; i + k <= hight; i ++ {
		for j := 0; j+k+i <= hight; j++ {
			subArr := nums[j:(i+k+j)]
			odd := 0
			for _ , v := range subArr {
				if v%2 == 1 {
					odd++
				}
				if odd == k {
					count++
					fmt.Println(subArr)
				}
			}
		}
	}
	return count
}
