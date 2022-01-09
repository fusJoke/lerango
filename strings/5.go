package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "完满 hello go 完满"
	str1 := " 请使用trimspace() 修剪空白字符   "

	fmt.Println(str)
	fmt.Printf("%q\n", strings.Trim(str, "完满"))
	fmt.Printf("%q\n", strings.Trim(str, "完满 "))
	fmt.Printf("%q\n", strings.TrimLeft(str, "完满 "))
	fmt.Printf("%q\n", strings.TrimRight(str, "完满 "))
	fmt.Printf("%q\n", strings.TrimSpace(str1))
}