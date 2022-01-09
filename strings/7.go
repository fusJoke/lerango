package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str  = "快乐(每)一天"
	var left = '('
	fmt.Println(utf8.RuneCountInString(str))
	for _, s := range str {
		if s == left {
			fmt.Println("find (")
		}
	}
}
