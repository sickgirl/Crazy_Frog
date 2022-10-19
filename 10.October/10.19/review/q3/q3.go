package main

import "fmt"

//1295. 统计位数为偶数的数字
//https://leetcode.cn/problems/find-numbers-with-even-number-of-digits/
func main() {
	nums := []int{12, 345, 2, 6, 7896}
	r := findNumbers(nums)
	fmt.Printf("%v", r)
}
func findNumbers(nums []int) int {
	r := 0
	for _, v := range nums {
		num := getNums(v)
		if num%2 == 0 {
			r++
		}
	}
	return r
}

func getNums(num int) int {
	r := 0
	for {
		if num < 10 {
			r++
			break
		}
		num = num / 10
		r++
	}
	return r
}
