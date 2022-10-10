package main

import (
	"fmt"
	"strings"
)

//2255. 统计是给定字符串前缀的字符串数目
//https://leetcode.cn/problems/count-prefixes-of-a-given-string/
func main() {
	words := []string{"e", "s", "mi", "e", "ia", "ibwu", "e", "e", "k", "ci", "rip", "suw", "a", "l"}
	s := "e"
	r := countPrefixes(words, s)
	fmt.Printf("%v", r)
}
func countPrefixes(words []string, s string) int {
	n := 0
	for _, v := range words {
		if strings.Index(s, v) == 0 {
			n++
		}
	}
	return n
}

//熟悉 strings 包的 index 用法
