package main

import (
	"fmt"
	"strings"
)

//反转字符串中的单词 III
//https://leetcode.cn/problems/reverse-words-in-a-string-iii/
func main() {
	s := "Let's take LeetCode contest"
	res := reverseWords(s)
	fmt.Printf("%v", res)
}

func reverseWords(s string) string {
	res1 := strings.Split(s, " ")
	var res []string
	for _, v := range res1 {
		temp := Reverse(v)
		res = append(res, temp)
	}
	return strings.Join(res, " ")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//解释: strings.spilt  相当于 php 中 的 explode
// strings.join 相当于 implode
