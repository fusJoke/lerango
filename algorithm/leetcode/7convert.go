package main

import (
	"bytes"
	"fmt"
)

/**
z字形变换
*/

func main() {
	var s string
	s = "abavabads"
	s = convert(s, 3)
	fmt.Print(s)
}

func convert(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}

	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}
