package main

import "fmt"

//执行操作后的变量值
//https://leetcode.cn/problems/final-value-of-variable-after-performing-operations/
func main() {
	ops := []string{
		"++X",
		"++X",
		"++X",
		"X++",
		"--X",
		"X--",
	}
	r := finalValueAfterOperations(ops)
	fmt.Printf("%v", r)
}
func finalValueAfterOperations(operations []string) int {
	r := 0
	for _, v := range operations {
		if v[1] == '+' {
			r++
		} else {
			r--
		}
	}
	return r
}
