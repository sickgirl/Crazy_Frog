package main

import (
	"fmt"
	"strings"
)

//1021. 删除最外层的括号
//https://leetcode.cn/problems/remove-outermost-parentheses/
func main() {
	s := "(()())(())"
	r := removeOuterParentheses(s)
	fmt.Printf("%v", r)
}
func removeOuterParentheses(s string) string {
	var strings2 []string
	tempN := 0
	tempS := ""
	lenN := 0
	for _, v := range s {
		tempS += string(v)
		lenN++
		if v == '(' {
			tempN++
		} else {
			tempN--
		}
		if tempN == 0 {
			strings2 = append(strings2, tempS[1:lenN-1])
			tempS = ""
			lenN = 0
		}
	}
	return strings.Join(strings2, "")
}

//可以优化 明天再说
