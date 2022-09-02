package main

//爬楼梯
//https://leetcode.cn/problems/climbing-stairs/
import "fmt"

func main() {
	n := climbStairs2(4)
	fmt.Printf("%v", n)
}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {
	var dp [46]int
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
