package main

//各位之积 - 各位之和
//https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
import "fmt"

func main() {
	n := 62
	r := subtractProductAndSum(n)
	fmt.Printf("%v", r)
}
func subtractProductAndSum(n int) int {
	x, y := 0, 1
	if n == 0 {
		return 0
	}
	for {
		if n == 0 {
			break
		}
		temp := n % 10

		x += temp
		y *= temp
		n = n / 10
	}

	return y - x
}

//解题思路:
//无需思路  拆成各个位的值 直接 加和 积和
