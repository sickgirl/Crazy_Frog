package main

import "fmt"

//数组串联
//https://leetcode.cn/problems/concatenation-of-array/
func main() {
	nums := []int{1, 2, 3, 4}
	r := getConcatenation(nums)
	fmt.Printf("%v", r)
}
func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
