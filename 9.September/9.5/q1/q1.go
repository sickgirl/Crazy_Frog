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
	if n == 0 {
		return 0
	}

	x, y := 0, 1
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
