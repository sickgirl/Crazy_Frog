package main

import "fmt"

//切片重组
func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	n := 5
	res := shuffle(nums, n)
	fmt.Printf("%v", res)
}

func shuffle(nums []int, n int) []int {
	r := make([]int, 0)
	for i := 0; i < n; i++ {
		r = append(r, nums[i], nums[i+n])
	}
	return r
}
