package main

import (
	"fmt"
	"math"
)

//剑指 Offer 17. 打印从1到最大的n位数
//https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func main() {
	num := 2
	r := printNumbers(num)
	fmt.Printf("%v", r)
}
func printNumbers(n int) []int {
	num := math.Pow(10, float64(n))
	num2 := int(num)
	r := []int{}
	for i := 1; i < num2; i++ {
		r = append(r, i)
	}
	return r
}
