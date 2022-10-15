package main

import "fmt"

//2176. 统计数组中相等且可以被整除的数对
//https://leetcode.cn/problems/count-equal-and-divisible-pairs-in-an-array/
func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	r := countPairs(nums, k)
	fmt.Printf("%v", r)
}
func countPairs(nums []int, k int) (ans int) {
	for j, y := range nums {
		for i, x := range nums[:j] {
			if x == y && i*j%k == 0 {
				ans++
			}
		}

	}
	return
}

//唯一需要注意的是   在嵌套循环遍历时候   外层为j  内层为i  因为最后要取ij的值  所以截取数组只能从后往前截
