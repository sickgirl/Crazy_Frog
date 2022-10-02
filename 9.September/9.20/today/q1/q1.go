package main

import "fmt"

//2413. 最小偶倍数
//https://leetcode.cn/problems/smallest-even-multiple/
func main() {
	n := 5
	r := smallestEvenMultiple(n)
	fmt.Printf("%v", r)
}
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
