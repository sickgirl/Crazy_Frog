package main

import "fmt"

//剑指 Offer 58 - II. 左旋转字符串
//https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func main() {
	s := "abcde"
	k := 2
	res := reverseLeftWords(s, k)
	fmt.Printf("%v", res)
}

func reverseLeftWords(s string, n int) string {
	s1 := s[:n]
	s2 := s[n:]
	return s2 + s1
}
