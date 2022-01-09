package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "this is a fat cat"
	fmt.Println("%t\n", strings.HasPrefix(str, "is"))

}
