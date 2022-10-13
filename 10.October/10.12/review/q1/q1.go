package main

import "fmt"

//1790. 仅执行一次字符串交换能否使两个字符串相等
//https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
func main() {
	s1 := "bank"
	s2 := "kanb"
	r := areAlmostEqual(s1, s2)
	fmt.Printf("%v", r)
}
func areAlmostEqual(s1, s2 string) bool {
	i, j := -1, -1
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			if i < 0 {
				i = idx
			} else if j < 0 {
				j = idx
			} else {
				return false
			}
		}
	}
	return i < 0 || j >= 0 && s1[i] == s2[j] && s1[j] == s2[i]
}
