package main

import "fmt"

//按奇偶排序数组 II
//https://leetcode.cn/problems/sort-array-by-parity-ii/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	r := sortArrayByParityII(nums)
	fmt.Printf("%v", r)
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	var j []int
	var o []int
	var r []int
	for _, v := range nums {
		if v%2 == 1 {
			j = append(j, v)
		} else {
			o = append(o, v)
		}
	}

	for i := 0; i < n/2; i++ {
		r = append(r, o[i])
		r = append(r, j[i])
	}
	return r
}
