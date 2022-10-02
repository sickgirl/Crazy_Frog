package main

import (
	"fmt"
	"strings"
)

//58. 最后一个单词的长度
//https://leetcode.cn/problems/length-of-last-word/
func main() {
	s := "Hello World"
	r := lengthOfLastWord(s)
	fmt.Printf("%v", r)
}

func lengthOfLastWord(s string) int {
	strArrayNew := strings.Split(s, " ")
	n := 0
	for _, v := range strArrayNew {
		if len(v) > 0 {
			n = len(v)
		}
	}
	return n
}
