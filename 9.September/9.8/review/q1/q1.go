package main

//括号的最大嵌套深度
//https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses/
import "fmt"

func main() {
	s := "(1)+((2))+(((3)))"
	max := maxDepth(s)
	fmt.Printf("%v", max)
}

func maxDepth(s string) int {
	max := 0
	len1 := 0
	for _, v := range s {
		if v == '(' {
			len1++
		}
		if v == ')' {
			len1--
		}
		if len1 >= max {
			max = len1
		}
	}
	return max
}
