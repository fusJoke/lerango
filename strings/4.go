package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world!\n"

	fmt.Printf("%s", str)
	fmt.Printf(strings.ToLower(str))
	fmt.Printf(strings.ToUpper(str))

}