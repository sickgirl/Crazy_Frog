package main

// 拆分数位后四位数字的最小和
// https://leetcode.cn/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
import (
	"fmt"
	"sort"
)

func main() {
	num := 1234 //结果是37
	r := minimumSum(num)
	fmt.Printf("%v", r)
}
func minimumSum(num int) int {
	var m []int
	for {
		if num == 0 {
			break
		}
		temp := 0
		temp = num % 10
		m = append(m, temp)
		num = num / 10
	}
	sort.Ints(m)
	return m[0]*10 + m[1]*10 + m[2] + m[3]
}
