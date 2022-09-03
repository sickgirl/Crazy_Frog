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
func mostWordsFound(sentences []string) int {
	max := 0
	for _, v1 := range sentences {
		temp := 0
		for _, v2 := range v1 {
			if string(v2) == " " {
				temp++
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max + 1
}

func mostWordsFound1(sentences []string) (ans int) {
	for _, s := range sentences {
		cnt := strings.Count(s, " ") + 1
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
