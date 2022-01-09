package main

import "fmt"

func main() {
	s := make([]int,5)
	s = append(s,1,2,3,4)
	fmt.Println(s)

	s2 := make([]int,0)
	s2 = append(s2,1,2,3,4)
	fmt.Println(s2)


}

/**
wangfusheng@MacBook-Pro-2 daily % go run 03day.go
[0 0 0 0 0 1 2 3 4]
[1 2 3 4]

func Mui(x, y int) (sum int, error) {
	sum =y
	return x+y,nil
}
 */