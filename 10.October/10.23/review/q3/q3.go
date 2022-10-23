package main

import (
	"fmt"
	"sort"
)

//1051. 高度检查器
//https://leetcode.cn/problems/height-checker/
func main() {
	nums := []int{1, 1, 4, 2, 1, 3}
	r := heightChecker(nums)
	fmt.Printf("%v", r)
}
func heightChecker(heights []int) int {
	test := append([]int{}, heights...)
	sort.Ints(test)
	r := 0
	for i, v := range heights {
		if test[i] != v {
			r++
		}
	}
	return r
}
