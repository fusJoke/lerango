package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "goalng is cool, right?"
	str1 := "科学建设钓鱼城特色的完满教育！"

	var manyG  string ="o"

	fmt.Printf("%d\n", strings.Count(str,manyG))
	fmt.Printf("%d\n", strings.Count(str,"o"))
	fmt.Printf("%d\n", strings.Count(str1, "完满教育"))
}
