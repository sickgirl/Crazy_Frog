package main

import (
	"fmt"
	"sort"
)

//1046. 最后一块石头的重量
//https://leetcode.cn/problems/last-stone-weight/
func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	r := lastStoneWeight(stones)
	fmt.Printf("%v", r)
}
func lastStoneWeight(stones []int) int {
	for {
		sort.Ints(stones)
		if len(stones) == 1 {
			return stones[0]
		} else if len(stones) == 0 {
			return 0
		}
		temp := stones[len(stones)-2:]
		stones = stones[:len(stones)-2]
		if temp[1] > temp[0] {
			stones = append(stones, temp[1]-temp[0])
		}
	}
	return 0
}
