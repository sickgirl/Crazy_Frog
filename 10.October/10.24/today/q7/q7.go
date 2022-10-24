package main

import "fmt"

//2220. 转换数字的最少位翻转次数
//https://leetcode.cn/problems/minimum-bit-flips-to-convert-number/
func main() {
	start := 10
	goal := 7
	r := minBitFlips(start, goal)
	fmt.Printf("%v", r)
}
func minBitFlips(start int, goal int) int {
	end := start ^ goal
	r := 0
	for {
		if end == 0 {
			return r
		}
		if end%2 == 1 {
			r++
		}
		end = end / 2
	}
}
