package main

//分糖果
//https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/

import "fmt"

func main() {
	candy := []int{1, 2, 3, 4, 5}
	extra := 3
	candies := kidsWithCandies(candy, extra)
	fmt.Printf("%v", candies)

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := GetMaxNum(candies)
	var r []bool
	for _, v := range candies {
		if v+extraCandies < max {
			r = append(r, false)
		} else {
			r = append(r, true)
		}
	}
	return r
}

func GetMaxNum(ary []int) int {
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal < ary[i] {
			maxVal = ary[i]
		}
	}
	return maxVal
}
