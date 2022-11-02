package main

import "fmt"

//1389. 按既定顺序创建目标数组
//https://leetcode.cn/problems/create-target-array-in-the-given-order/
func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	r := createTargetArray(nums, index)
	fmt.Printf("%v", r)
}

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i, v := range index {
		if v == len(res) {
			res = append(res, nums[i])
		} else {
			t1 := res[:v]
			t2 := res[v:]
			t3 := append([]int{nums[i]}, t2...)
			res = append(t1, t3...)
		}
	}
	return res
}
