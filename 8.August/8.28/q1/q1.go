package main

import "fmt"

//两数之和
//https://leetcode.cn/problems/two-sum/

func main() {
	res := twoSum2([]int{2, 7, 4, 4}, 9)

	fmt.Printf("%v", res)
}

func twoSum(nums []int, target int) []int {
	for r1, v := range nums {
		temp_target := target - v
		for r2, v1 := range nums {
			if (temp_target == v1) && (r1 != r2) {
				return []int{r1, r2}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	map1 := make(map[int]int)

	for r, v := range nums {
		fmt.Println("第", r+1, "轮:")
		if _, ok := map1[target-v]; ok {
			fmt.Println("命中了:")

			return []int{map1[target-v], r}
		} else {
			fmt.Println("没有命中:")
			map1[v] = r

			fmt.Printf("当前map1:  %v", map1)
		}

	}

	return []int{}
}
