package main

//面试题 01.09. 字符串轮转
//https://leetcode.cn/problems/string-rotation-lcci/
import (
	"fmt"
	"strings"
)

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	r := isFlipedString(s1, s2)
	fmt.Printf("%v", r)
}

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
