package main

import (
	"fmt"
	"math"
)

//LCP 06. 拿硬币
//https://leetcode.cn/problems/na-ying-bi/
func main() {
	nums := []int{4, 2, 1}
	r := minCount(nums)
	fmt.Printf("%v", r)
}

func minCount(coins []int) int {
	n := 0
	for _, v := range coins {
		temp := float64(v) / 2
		temp = math.Ceil(temp)
		n += int(temp)
	}
	return n
}
