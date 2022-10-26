package main

import "fmt"

//1486. 数组异或操作
//https://leetcode.cn/problems/xor-operation-in-an-array/
func main() {
	n := 4
	start := 3
	r := xorOperation(n, start)
	fmt.Printf("%v", r)
}
func xorOperation(n int, start int) int {
	r := 0
	for i := start; i < 2*n+start; i += 2 {
		r ^= i
	}
	return r
}
