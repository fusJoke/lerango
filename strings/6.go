package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is golang Project"
	fmt.Println(str)
	fmt.Printf("%q\n", strings.Split(str, "Project"))
	fmt.Printf("%q\n", strings.Split(str, " "))
}
