package main

import "fmt"

//猜数字
func main() {
	r := game([]int{1, 2, 3}, []int{3, 3, 3})
	fmt.Printf("%v", r)
}
func game(guess []int, answer []int) int {
	n := 0
	for i, v := range guess {
		if v < answer[i] {
			n++
		}
	}
	return n
}
