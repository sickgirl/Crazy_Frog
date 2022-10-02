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
	r1 := arrayStringsAreEqual1(a, b)
	fmt.Printf("%v  ", r1)
}

func arrayStringsAreEqual1(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

//strings.join    相当于php 的implode
//strings.FieldsFunc(srcStr2, splitFunc)   相当于 php 中的 explode
