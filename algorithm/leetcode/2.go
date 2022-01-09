package main

import "fmt"

func main() {
	var str string = "lee(t(c)o)de)"
	fmt.Println(minRemoveToMakeValid(str))
}

/*
给你一个由 '('、')' 和小写字母组成的字符串 s。

你需要从字符串中删除最少数目的 '(' 或者 ')'（可以删除任意位置的括号)，使得剩下的「括号字符串」有效。

请返回任意一个合法字符串。

有效「括号字符串」应当符合以下任意一条要求：

空字符串或只包含小写字母的字符串
可以被写作AB（A连接B）的字符串，其中A和B都是有效「括号字符串」
可以被写作(A)的字符串，其中A是一个有效的「括号字符串」

 */

func minRemoveToMakeValid(s string) string {


	for _, v := range s {
		if v == '('  {

		}
		if v == ')' {

		}
	}
	return s
}