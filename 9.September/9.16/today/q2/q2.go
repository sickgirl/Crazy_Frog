package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 1, 1, 3}
	r := numIdenticalPairs(test)
	fmt.Printf("%v", r)
}
func numIdenticalPairs(nums []int) int {
	sameNum := make(map[int]int)
	for _, v := range nums {
		sameNum[v]++
	}
	var result int
	for _, v := range sameNum {
		result += (v - 1) * v / 2
	}
	return result
}
