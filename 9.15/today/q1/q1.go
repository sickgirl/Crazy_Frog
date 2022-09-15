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
	res := ""
	temp := ""
	for _, v := range s {
		if v != ' ' {
			temp = string(v) + temp
		} else {
			res = res + temp + " "
			temp = ""
		}
	}
	res = res + temp
	return res
}
