package main

import "fmt"

//541. 反转字符串 II
//https://leetcode.cn/problems/reverse-string-ii/
func main() {
	s := "abcdefg"
	k := 2
	r := reverseStr(s, k)
	fmt.Printf("%v", r)
}

func reverseStr(s string, k int) string {
	res := ""
	for {
		n := len(s)
		if n < k {
			res = res + Bytes(s)
			break
		} else if n < 2*k {
			res = res + Bytes(s[:k]) + s[k:]
			break
		}

		tempS := s[:k]
		tempS = Bytes(tempS)
		temp2 := s[k : 2*k]
		res = res + tempS + temp2
		s = s[2*k:]
	}
	return res
}

func Bytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-i-1]
	}
	return string(r)
}
