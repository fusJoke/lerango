package main

import "fmt"

/**
结构体的比较
 */
func main() {
	sn1 := struct {
		age int
		name string
	}{age:10, name:"阿花"}
	sn2 := struct {
		age int
		name string
	}{age:10, name:"阿花"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age:10,m:map[string]string{"a":"1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	/**
		map是不可比较的，所以两个结构体就是不可比较的。编译的时候会出问题
	 */
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}