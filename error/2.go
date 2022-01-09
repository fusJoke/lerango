package main

import (
	"errors"
	"fmt"
)
// 索引越界， 不可恢复的环境问题，栈溢出采用panic
func main() {
	Check(1);
}

func Positive(n int) (bool, error){
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n>-1, nil
}

func Check(n int) {
	pos, err := Positive(n)
	if err != nil {
		fmt.Println(n, err)
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}
