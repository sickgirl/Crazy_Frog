package main

import "fmt"

//1512. 好数对的数目
//https://leetcode.cn/problems/number-of-good-pairs/
func main() {
	test := []int{1, 2, 3, 1, 1, 3}
	r := numIdenticalPairs(test)
	fmt.Printf("%v", r)
}
func numIdenticalPairs(nums []int) int {
	sameNum := make(map[int]int)
	for _, v := range nums {
		sameNum[v]++
	}
	var result int
	for _, v := range sameNum {
		result += (v - 1) * v / 2
	}
	return result
}

//题解: 答案只和 每个值的个数有关   比如 1有3个   那对应的号数对组合  就是  c3,2 =   3 * (3-1) /2 = 3
//每个值 有n个  则组合为  cn,2   = n(n-1)/2
