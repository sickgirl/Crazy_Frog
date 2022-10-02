package main

import (
	"fmt"
	"strings"
)

//1662. 检查两个字符串数组是否相等
//https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/
func main() {
	a := []string{"a", "bc"}
	b := []string{"ab", "c"}
	r := arrayStringsAreEqual(a, b)
	r1 := arrayStringsAreEqual1(a, b)
	fmt.Printf("%v  %v", r, r1)
}

func arrayStringsAreEqual1(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	la := getLe(word1)
	lb := getLe(word2)
	return la == lb
}
func getLe(a []string) string {
	r := ""
	for _, v := range a {
		r += v
	}
	return r
}
