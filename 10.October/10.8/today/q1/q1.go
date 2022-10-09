package main

import "fmt"

//1588. 所有奇数长度子数组的和
//https://leetcode.cn/problems/sum-of-all-odd-length-subarrays/
func main() {
	nums := []int{1, 4, 2, 5, 3}
	r := sumOddLengthSubarrays(nums)
	fmt.Printf("%v", r)
}

func sumOddLengthSubarrays(arr []int) (sum int) {
	n := len(arr)

	for start := range arr {
		for length := 1; start+length <= n; length += 2 {
			for _, v := range arr[start : start+length] {
				sum += v
			}
		}
	}
	return sum
}
