package main

import "fmt"

//两数之和

func main() {
	nums := []int{2, 7, 13, 4}
	target := 13

	res := threeSum(nums, target)
	fmt.Printf("%v", res)
}

func towSum(nums []int, target int) []int {
	map1 := make(map[int]int)
	for r, v := range nums {
		if _, ok := map1[target-v]; ok {
			return []int{map1[target-v], r}
		} else {
			map1[v] = r
		}
	}
	return []int{}
}

func threeSum(nums []int, target int) []int {
	for r1, v1 := range nums {
		for r2, v2 := range nums {
			for r3, v3 := range nums {
				if (v1+v2+v3 == target) && (r1 != r2) && (r1 != r3) {
					return []int{r1, r2, r3}
				}
			}
		}
	}
	return []int{}

}
