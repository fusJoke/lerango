package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string = "this is a 字符串"
	var length int
	length = utf8.RuneCountInString(str)
	fmt.Println(length)
	var by []byte = []byte(str)
	res := utf8.FullRune(by)
	fmt.Println(res)
}