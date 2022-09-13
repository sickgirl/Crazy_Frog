package main

//去除某个元素
//https://leetcode.cn/problems/remove-element/

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 2
	r := removeElement(nums, val)
	fmt.Printf("%v", r)
}
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)

}
