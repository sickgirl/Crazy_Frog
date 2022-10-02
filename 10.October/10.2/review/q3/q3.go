package main

import "fmt"

//剑指 Offer 39. 数组中出现次数超过一半的数字
//https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
func main() {
	number := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	r := majorityElement(number)
	fmt.Printf("%v", r)
}

func majorityElement(nums []int) int {
	ans := map[int]int{}
	n := len(nums) / 2
	for _, v := range nums {
		ans[v]++
		if ans[v] > n {
			return v
		}
	}
	return 0
}
