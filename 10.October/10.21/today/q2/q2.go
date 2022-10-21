package main

import "fmt"

//509. 斐波那契数
//https://leetcode.cn/problems/fibonacci-number/
func main() {
	n := 10
	r := fib(n)
	fmt.Printf("%v", r)

}
func fib(n int) int {
	ans := map[int]int{
		0: 0,
		1: 1,
	}
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			continue
		}
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans[n]
}
