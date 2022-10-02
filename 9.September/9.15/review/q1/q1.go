package main

import "fmt"

//基于排列构建数组
//https://leetcode.cn/problems/build-array-from-permutation/
func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	ans := buildArray(nums)
	println(ans)
	fmt.Printf("%v", ans)
}

func buildArray(nums []int) []int {
	n := len(nums)
	res := []int{}
	for i := 0; i <= n-1; i++ {
		res = append(res, nums[nums[i]])
	}
	return res
}
