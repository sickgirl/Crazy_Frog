package main

import (
	"fmt"
	"math"
)

//求是否为2的幂
//https://leetcode.cn/problems/power-of-two/
func main() {
	n := 222
	r := isPowerOfTwo(n)
	fmt.Printf("%v", r)
}
func isPowerOfTwo(n int) bool {
	i := 0
	for {
		temp := math.Pow(2, float64(i))
		if int(temp) == n {
			return true
		}
		if int(temp) > n {
			return false
		}
		i++
	}
	return true
}

//题解思路  每次乘2  看是否超过目标值了
