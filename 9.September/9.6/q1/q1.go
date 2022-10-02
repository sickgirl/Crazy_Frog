package main

//速算机器人
//https://leetcode.cn/problems/nGK0Fy/
import (
	"fmt"
	"math"
)

func main() {
	c := "AB"
	r := calculate1(c)
	fmt.Printf("%v", r)
}

func calculate(s string) int {
	x, y := 1, 0
	for _, s1 := range s {
		if string(s1) == "A" {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}
func calculate1(s string) int {
	n := len(s)
	r := math.Pow(2, float64(n))
	return int(r)
}

//解题思路:
//不须遍历直接计算的话  可以拆开看 A  x = 2x + y   这时候计算 总值   y不变  x = x + (x+y) 即  x+y = 2(x+y)
//不须遍历直接计算的话  可以拆开看 B  y = x + 2y   这时候计算 总值   x不变  y = y + (x+y) 即  x+y = 2(x+y)
//所以 AB 没差别  只要看 操作式长度   即可 1 * 2^len(c)
