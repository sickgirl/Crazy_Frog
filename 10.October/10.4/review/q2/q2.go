package main

import "fmt"

//1863. 找出所有子集的异或总和再求和
//https://leetcode.cn/problems/sum-of-all-subset-xor-totals/
func main() {
	nums := []int{1, 3}
	r := subsetXORSum(nums)
	fmt.Printf("%v", r)
}
func subsetXORSum(nums []int) (ans int) {
	var dfs func(i int, cur int)
	dfs = func(i int, cur int) {
		if i == len(nums) {
			return
		}
		dfs(i+1, cur)  //跳过当前数
		cur ^= nums[i] // 增加当前数
		ans += cur
		dfs(i+1, cur)
	}
	dfs(0, 0)
	return
}
