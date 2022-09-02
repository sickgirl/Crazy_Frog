package main

//切片重排
//https://leetcode.cn/problems/shuffle-the-array/
import "fmt"

func main() {

	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	n := 5
	res := shuffle(nums, n)
	fmt.Printf("%v", res)
}
func shuffle(nums []int, n int) []int {
	s1 := make([]int, 0)
	for i := 0; i < n; i++ {
		s1 = append(s1, nums[i], nums[i+n])
	}
	return s1
}
