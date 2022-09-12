package main

//删除重复数据
//https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
import (
	"fmt"
)

func main() {
	r := removeDuplicates([]int{1, 1, 2, 2, 3, 4})
	fmt.Print(r)
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
