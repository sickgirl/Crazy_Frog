package main

import "fmt"

//两数之和
//https://leetcode.cn/problems/two-sum/

func main() {
	res := twoSum2([]int{2, 7, 4, 4}, 9)

	fmt.Printf("%v", res)
}

func twoSum2(nums []int, target int) []int {
	map1 := make(map[int]int)

	for r, v := range nums {
		tempRes := target - v
		if _, ok := map1[tempRes]; ok {
			return []int{map1[tempRes], r}
		}
		map1[v] = r
	}

	return []int{}
}
