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
	len := 0
	for _, v := range s {
		temp := string(v)
		if temp == "(" {
			len++
		}
		if temp == ")" {
			len--
		}
		if len > max {
			max = len
		}
	}
	return max
}
