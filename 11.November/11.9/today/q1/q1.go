package main

import (
	"fmt"
	"strings"
)

//1374. 生成每种字符都是奇数个的字符串
//https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts/
func main() {
	n := 4
	r := generateTheString(n)
	fmt.Printf("%v", r)
}
func generateTheString(n int) string {
	if n%2 == 0 {
		//偶数
		x1 := strings.Repeat("a", n-1)
		return x1 + "b"

	} else {
		if n == 1 {
			return "a"
		}
		x1 := strings.Repeat("a", n-1)
		return x1 + "b" + "c"
	}
}
