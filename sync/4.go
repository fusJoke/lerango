package main

import (
	"fmt"
)

var val int
var val2 *int

func main() {
	fmt.Println(val)
	var mapx map[string]int = make(map[string]int)
	mapx["name"] = 1
	fmt.Println(mapx)
}