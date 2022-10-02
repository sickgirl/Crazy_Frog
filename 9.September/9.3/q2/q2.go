package main

import "fmt"

//计算比赛场次
//https://leetcode.cn/problems/count-of-matches-in-tournament/
func main() {
	n := 11
	r := numberOfMatches(n)
	fmt.Printf("%v", r)
}
func numberOfMatches(n int) int {
	return n - 1 //题解分析   每场淘汰一个队伍  最终胜一个   所以需要比赛n-1
}
