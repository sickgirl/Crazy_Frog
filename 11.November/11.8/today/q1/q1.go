package main

import (
	"fmt"
	"strings"
)

//剑指 Offer 05. 替换空格
//https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func main() {
	s := "We are happy."
	r := replaceSpace(s)
	fmt.Printf("%v", r)
}
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
