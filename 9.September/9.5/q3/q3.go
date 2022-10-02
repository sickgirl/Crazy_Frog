package main

import (
	"fmt"
	"strings"
)

//分割字符串的最大得分
//https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
func main() {
	s := "10011"
	r := maxScore(s)
	fmt.Printf("%v", r)
}

func maxScore(s string) int {
	full := len(s)
	if full == 1 {
		return 1
	}
	if full == 0 {
		return 0
	}
	max := 0
	for i := 1; i <= full-1; i++ {

		left := s[:i]
		right := s[i:]
		//fmt.Printf("%v %v %v", i, left, right)
		no1 := strings.Count(left, "0")
		no2 := strings.Count(right, "1")
		if no2+no1 > max {
			max = no1 + no2
		}
	}
	return max
}
