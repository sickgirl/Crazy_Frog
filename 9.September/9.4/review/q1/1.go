package main

import (
	"fmt"
	"strings"
)

//查找单词最多的句子的单词数
//https://leetcode.cn/problems/maximum-number-of-words-found-in-sentences
func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	r := mostWordsFound(sentences)
	fmt.Printf("%v", r)
}

func mostWordsFound(sentences []string) (max int) {
	max = 0
	for _, v := range sentences {
		one := strings.Count(v, " ")
		if one > max {
			max = one
		}
	}
	return max + 1
}
