package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "撒大声地123sad撒232挨过打发生"
	fmt.Printf("%d\n", len([]rune(str)))
	fmt.Println(utf8.RuneCountInString(str))
}