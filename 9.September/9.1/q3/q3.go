package main

import "fmt"

//猜数字
//https://leetcode.cn/problems/guess-numbers/
func main() {
	r := game([]int{1, 2, 3}, []int{3, 3, 3})
	fmt.Printf("%v", r)
}

func game(guess []int, answer []int) int {
	r := 0
	for i, v := range guess {
		if v == answer[i] {
			r++
		}
	}
	return r
}
