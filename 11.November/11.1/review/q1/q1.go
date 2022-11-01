package main

import (
	"fmt"
	"strings"
)

//709. 转换成小写字母
//https://leetcode.cn/problems/to-lower-case/
func main() {
	s := "Hello"
	r := toLowerCase(s)
	fmt.Printf("%v", r)
}
func toLowerCase(s string) string {
	return strings.ToLower(s)
}
