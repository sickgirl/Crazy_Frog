package main

import "fmt"

//2367. 算术三元组的数目
//https://leetcode.cn/problems/number-of-arithmetic-triplets/
func main() {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	r := arithmeticTriplets(nums, diff)
	fmt.Printf("%v", r)
}
func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	ans := map[int]bool{}
	res := 0
	for _, v := range nums {
		ans[v] = true
	}

	for i := 0; i < n-2; i++ {
		r := getBool(nums[i], diff, ans)
		if r {
			res++
		}
	}
	return res
}

func getBool(num int, diff int, ans map[int]bool) bool {
	start := num
	for i := 1; i < 3; i++ {
		start += diff
		if !ans[start] {
			return false
		}
	}
	return true
}

//其实就是从 一个有序数组中 寻找 差为 diff 的  长度为3的等差数列
