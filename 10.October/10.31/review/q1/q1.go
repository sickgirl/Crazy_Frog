package main

import "fmt"

//349. 两个数组的交集
//https://leetcode.cn/problems/intersection-of-two-arrays/
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	r := intersection(nums1, nums2)
	fmt.Printf("%v", r)
}

func intersection(nums1 []int, nums2 []int) []int {
	temp1 := map[int]bool{}
	r := map[int]bool{}
	for _, v := range nums1 {
		temp1[v] = true
	}
	for _, v := range nums2 {
		if temp1[v] {
			r[v] = true
		}
	}
	var res []int
	for i, _ := range r {
		res = append(res, i)
	}
	return res
}
