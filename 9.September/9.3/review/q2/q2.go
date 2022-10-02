package main

import "fmt"

//只出现一次的数字
//https://leetcode.cn/problems/single-number/

func main() {
	nums := []int{2, 2, 1, 3, 3}
	r := singleNumber(nums)
	fmt.Print(r)
}
func singleNumber(nums []int) int {
	r := 0
	for _, v := range nums {
		r = r ^ v
	}
	return r
}
