package main

//分割平衡字符串
//https://leetcode.cn/problems/split-a-string-in-balanced-strings/
import "fmt"

func main() {
	s := "RLRL"
	r := balancedStringSplit(s)
	fmt.Printf("%v", r)
}

func balancedStringSplit(s string) int {
	l := 0
	res := 0
	for _, v := range s {
		if v == 'L' {
			l++
		} else {
			l--
		}
		if l == 0 {
			res++
		}
	}
	return res
}
