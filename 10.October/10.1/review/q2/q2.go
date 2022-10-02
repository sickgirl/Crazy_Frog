package main

import "fmt"

//1480. 一维数组的动态和
//https://leetcode.cn/problems/running-sum-of-1d-array/
func main() {
	nums := []int{1, 2, 3, 4}
	r := runningSum(nums)
	fmt.Printf("%v", r)
}
func runningSum(nums []int) []int {
	n := 0
	var res []int
	for _, v := range nums {
		n += v
		res = append(res, n)

	}
	return res
}
