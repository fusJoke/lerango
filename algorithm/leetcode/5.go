package main

import "fmt"

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
//

func main() {
	fmt.Println(isPalindrome(123321))
}

func isPalindrome(x int) bool {
	var rev int = 0
	cp := x
	for x != 0 {
		digit := x % 10
		x /= 10
		rev = rev * 10 + digit
	}
	fmt.Println(rev)
	return rev == cp
}