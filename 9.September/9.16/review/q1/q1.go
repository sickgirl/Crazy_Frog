package main

import "fmt"

//反转字符串中的单词 III
//https://leetcode.cn/problems/reverse-words-in-a-string-iii/
func main() {
	s := "Let's take LeetCode contest"
	res := reverseWords(s)
	fmt.Printf("%v", res)
}

func reverseWords(s string) string {
	var res []byte
	var temp []byte
	for _, v := range s {
		if v != ' ' {
			temp = append([]byte{byte(v)}, temp...)
		} else {
			res = append(res, temp...)
			res = append(res, ' ')
			temp = []byte{}
		}
	}
	res = append(res, temp...)
	return string(res)
}
