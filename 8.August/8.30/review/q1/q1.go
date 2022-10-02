package main

import "fmt"

//爬楼梯
func main() {
	n := climbStairs2(4)
	fmt.Printf("%v", n)
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make(map[int]int)
	if n >= 3 {

		dp[1] = 1
		dp[2] = 2
		for i := 3; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}

	}
	return dp[n]

}
