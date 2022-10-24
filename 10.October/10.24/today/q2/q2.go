package main

import "fmt"

//1025. 除数博弈
//https://leetcode.cn/problems/divisor-game/
func main() {
	n := 2
	r := divisorGame(n)
	fmt.Printf("%v", r)
}
func divisorGame(n int) bool {
	return n%2 == 0
}
