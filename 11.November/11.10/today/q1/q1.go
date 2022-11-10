package main

import "fmt"

//LCP 07. 传递信息
//https://leetcode.cn/problems/chuan-di-xin-xi/
func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	r := findFinalValue(nums, original)
	fmt.Printf("%v", r)
}
func findFinalValue(nums []int, original int) int {

	map1 := map[int]bool{}
	for _, v := range nums {
		map1[v] = true
	}

	for {
		_, ok := map1[original]
		if ok {
			original *= 2
		} else {
			break
		}
	}
	return original
}
