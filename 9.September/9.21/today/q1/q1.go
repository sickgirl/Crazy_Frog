package main

import (
	"fmt"
	"math"
)

//2006. 差的绝对值为 K 的数对数目
//https://leetcode.cn/problems/count-number-of-pairs-with-absolute-difference-k/
func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	r := countKDifference(nums, k)
	fmt.Printf("%v", r)

}
func countKDifference(nums []int, k int) int {
	r := 0
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			temp := v - nums[j]
			if math.Abs(float64(temp)) == float64(k) {
				r++
			}
		}
	}
	return r
}

func countKDifference1(nums []int, k int) (ans int) {
	v := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if c, ok := v[nums[i]+k]; ok {
			ans += c
		}
		if c, ok := v[nums[i]-k]; ok {
			ans += c
		}
		v[nums[i]]++
	}
	return
}
