package main

import "fmt"

//1800. 最大升序子数组和
//https://leetcode.cn/problems/maximum-ascending-subarray-sum/
func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	res := maxAscendingSum(nums)
	fmt.Printf("%v", res)
}
func maxAscendingSum(nums []int) int {
	var res []int
	max := 0
	for i, v := range nums {
		temp := v
		tempList := nums[i:]
		sum := 0
		for i, v1 := range tempList {
			if i == 0 {
				sum += v1
				continue
			}
			if v1 > temp {
				sum += v1
				temp = v1
			} else {
				break
			}
		}
		res = append(res, sum)
	}

	for _, v := range res {
		if v > max {
			max = v
		}
	}

	return max
}
