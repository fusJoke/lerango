package main

import "fmt"

type Myint1 int
type Myint2 = int
/**
	Myint1 是基于int创建的新类型
	Myint2 是类型的别名
 */
func main() {
	var i int =0
	var i1 Myint1 = i	//cannot use i (type int) as type MyInt1 in assignment
	var i2 Myint2 = i
	fmt.Println(i1,i2)

}
