package main

import "fmt"

//汉明重量
//https://leetcode.cn/problems/number-of-1-bits/
func main() {
	var num uint32 = 00000000000000000000000000001011
	r := hammingWeight(num)
	fmt.Printf("%v", r)
}

func hammingWeight(num uint32) int {
	r := 0
	for {
		if num == 0 {
			return r
		}
		yu := num % 2
		if yu == 1 {
			r += 1
		}

		println(yu, ":", num, ":", r)
		num = num / 2
	}
}
