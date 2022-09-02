package main

//错误的集合
//https://leetcode.cn/problems/set-mismatch/

import "fmt"

//判断是否为回文

func main() {
	nums := []int{
		1, 2, 3, 4, 6, 7, 7, 8,
	}
	r := findErrorNums(nums)
	fmt.Printf("%v", r)
}
func findErrorNums(nums []int) []int {
	miss := 0
	double := 0

	have := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := have[nums[i]]; ok {
			double = nums[i]
		}
		have[nums[i]] = 1
		if i == 0 && nums[i] != 1 {
			miss = 1
		}
		if (i != 0) && ((nums[i] - nums[i-1]) == 2) {
			miss = nums[i] - 1
		}
	}
	if miss == 0 {
		miss = nums[len(nums)-1] + 1
	}
	return []int{double, miss}
}
