package main

//496. 下一个更大元素 I
//https://leetcode.cn/problems/next-greater-element-i/
func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 3, 4}
	r := nextGreaterElement(nums1, nums2)
	println(r)
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	r := []int{}
	for _, v := range nums1 {
		start := false
		write := false
		for _, v1 := range nums2 {
			if v1 == v {
				start = true
			}
			if start && v1 > v {
				r = append(r, v1)
				write = true
				break
			}
		}
		if !write {
			r = append(r, -1)
		}

	}
	return r
}
