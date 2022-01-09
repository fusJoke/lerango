package main

import "C"
import "fmt"

const(
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)
/**
0 2 zz zz 5
 */

func main() {
	fmt.Println(x,y,z,k,p)

	var x = nil				//A
	var x interface{} = nil //B
	var x string = nil		//C
	var x error = nil		//D
	/*
	BD 正确；nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
	 */
}

