package main

import "fmt"

//最大连续 1 的个数
//https://leetcode.cn/problems/max-consecutive-ones/
func main() {
	s := []int{1, 1, 0, 1, 1, 1}
	max := findMaxConsecutiveOnes(s)
	fmt.Printf("%v", max)
}

func findMaxConsecutiveOnes(nums []int) int {
	temp := 0
	max := 0
	for _, v := range nums {
		if v == 0 {
			if temp > max {
				max = temp
				temp = 0
			}
		} else {
			temp++
		}
	}
	if temp > max {
		max = temp
		temp = 0
	}
	return max
}

//复习到 9.14
