package main

import (
	"fmt"
	"strings"
)

//1784. 检查二进制字符串字段
//https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
func main() {
	s := "100"
	r := checkOnesSegment(s)
	fmt.Printf("%v", r)
}

func checkOnesSegment(s string) bool {
	if strings.Count(s, "01") > 0 {
		return false
	}
	return true
}
