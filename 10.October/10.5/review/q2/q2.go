package main

import (
	"fmt"
	"sort"
)

//2089. 找出数组排序后的目标下标
//https://leetcode.cn/problems/find-target-indices-after-sorting-array/
func main() {
	numbers := []int{1, 2, 5, 2, 3}
	target := 2
	r := targetIndices(numbers, target)
	fmt.Printf("%v", r)
}

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	r := []int{}
	for i, v := range nums {
		if v == target {
			r = append(r, i)
		} else if v > target {
			break
		}
	}
	return r
}
