package main

import "fmt"

//1748. 唯一元素的和
//https://leetcode.cn/problems/sum-of-unique-elements/
func main() {
	nums := []int{1, 2, 3, 2}
	r := sumOfUnique(nums)
	fmt.Printf("%v", r)
}
func sumOfUnique(nums []int) int {
	resMap := map[int]int{}
	res := 0
	for _, v := range nums {
		resMap[v]++
		if resMap[v] == 1 {
			res += v
		} else if resMap[v] == 2 {
			res -= v
		}

	}
	return res
}
