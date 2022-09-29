package main

import "fmt"

//2278. 字母在字符串中的百分比
//https://leetcode.cn/problems/percentage-of-letter-in-string/
func main() {
	s := "foobar"
	letter := 'o'
	r := percentageLetter(s, byte(letter))
	fmt.Printf("%v", r)
}
func percentageLetter(s string, letter byte) int {
	total := len(s)
	n := 0
	for _, v := range s {
		if byte(v) == letter {
			n++
		}
	}
	return n * 100 / total
}
