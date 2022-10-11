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
	for i := 0; i < n; i++ {
		temp := start + 2*i
		r ^= temp
	}
	return r
}
