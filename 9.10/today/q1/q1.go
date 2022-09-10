package main

import "fmt"

//子集异或之和
//https://leetcode.cn/problems/sum-of-all-subset-xor-totals/
func main() {
	nums := []int{1, 3}
	r := subsetXORSum(nums)
	fmt.Printf("%v", r)
}

func subsetXORSum(nums []int) int {
	var res int
	var tmp []int
	var dfs func(index int)
	dfs = func(index int) {
		var sum int
		for _, v := range tmp {
			sum ^= v
		}
		res += sum
		for i := index; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)
	return res
}

func subsetXORSum1(nums []int) int {

	result := 0

	var dfs func(sum int, i int)

	dfs = func(sum, i int) {
		if i >= len(nums) {
			result += sum
			return
		}
		dfs(sum, i+1)
		dfs(sum^nums[i], i+1)
	}

	dfs(0, 0)

	return result
}
