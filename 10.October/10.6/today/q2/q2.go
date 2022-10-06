package main

import (
	"fmt"
	"sort"
)

//561. 数组拆分
//https://leetcode.cn/problems/array-partition/
func main() {
	num := []int{1, 4, 3, 2}
	r := arrayPairSum(num)
	fmt.Printf("%v", r)
}

func arrayPairSum(nums []int) int {
	n := len(nums)
	r := 0
	sort.Ints(nums)
	for i := 0; i < n; i += 2 {
		r += nums[i]
	}
	return r
}
