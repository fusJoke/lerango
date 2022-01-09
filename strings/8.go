package main

import (
	"fmt"
	"strings"
)

//比较字符串

func main() {
	var s1 string = "this is normal string"
	var s2 string = "this is normal string"
	var s3 string = "this is Normal String"

	if s1 == s2 {
		fmt.Println(111)
	}

	if strings.Compare(s1, s3) == 0 {
		fmt.Println(222)
	}

	if strings.EqualFold(s1, s3) {
		fmt.Println(333)
	}


}