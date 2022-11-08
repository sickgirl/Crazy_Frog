package main

import (
	"fmt"
	"strings"
)

//1668. 最大重复子字符串
//https://leetcode.cn/problems/maximum-repeating-substring/
func main() {
	sequence := "ababc"
	word := "ab"
	r := maxRepeating(sequence, word)
	fmt.Printf("%v", r)
}
func maxRepeating(sequence string, word string) int {
	for k := len(sequence) / len(word); k > 0; k-- {
		if strings.Contains(sequence, strings.Repeat(word, k)) {
			return k
		}
	}
	return 0
}
