package main

import "fmt"

//1822. 数组元素积的符号
//https://leetcode.cn/problems/sign-of-the-product-of-an-array/
func main() {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	res := arraySign(nums)
	fmt.Printf("%v", res)
}

func arraySign(nums []int) int {
	r := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			r = -r
		}
	}
	return r
}
