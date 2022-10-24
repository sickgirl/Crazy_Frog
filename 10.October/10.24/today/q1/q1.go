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
	n := len(nums)
	var res []int
	for i := 0; i < n; i++ {
		tempList := nums[i:]
		tempSum := 0
		curent := tempList[0]
		for i, v := range tempList {
			if i == 0 {
				tempSum += v
				continue
			}
			if v <= curent {
				break
			}
			if v > curent {
				curent = v
				tempSum += v
			}

		}
		res = append(res, tempSum)
	}
	max := 0
	for _, v := range res {
		if v > max {
			max = v
		}
	}
	return max
}
