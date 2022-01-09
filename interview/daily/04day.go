package main

import "fmt"

/**
var(
    size := 1024
    max_size = size*2
)
不能编译通过，只能在函数内部使用简短模式；
 */

func main() {
	//list := new([]int) new 返回的是一个指针，无法对指针进行append操作。
	list := make([]int,0)
	list = append(list, 1)
	fmt.Println(list)

	/**
	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s2 = append(s2,s1)
	第二次参数不能使用slice，应该使用s1...操作符,将一个切片追加到另一个切片上
	 */
	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s2 = append(s2,s1...)

}
