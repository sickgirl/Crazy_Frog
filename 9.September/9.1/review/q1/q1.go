package main

import (
	"fmt"
	"sort"
)

//错误的集合
func main() {
	nums := []int{
		1, 2, 3, 4, 6, 7, 7, 8,
	}
	r := findErrorNums(nums)
	fmt.Printf("%v", r)
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	temp := 0
	miss := 0
	double := 0
	for _, v := range nums {
		if temp == v {
			double = v
		}
		if v-temp == 2 {
			miss = v - 1
		}
		temp = v
	}
	if miss == 0 {
		miss = nums[len(nums)-1] + 1
	}
	return []int{double, miss}
}
