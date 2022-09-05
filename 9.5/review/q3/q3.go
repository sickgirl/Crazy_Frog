package main

//分糖果
//https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/

import (
	"fmt"
	"sort"
)

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
	sort.Sort(sort.Reverse(sort.IntSlice(ary)))
	return ary[0]
}

//解题思考:
//当把全部额外糖都给他一个人时  他都比不上第一 (或者说成不了第一) 那他就没希望了
