package main

import "fmt"

//496. 下一个更大元素 I
//https://leetcode.cn/problems/next-greater-element-i/
func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 3, 4}
	r := nextGreaterElement(nums1, nums2)
	println(r)
	fmt.Printf("%v", r)
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int
	for _, v := range nums1 {
		start := false
		max := -1
		for _, t := range nums2 {
			if v == t {
				start = true
			}
			if start && t > v {
				max = t
				break
			}
		}
		res = append(res, max)

	}
	return res
}
