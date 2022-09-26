package main

import "fmt"

//1365. 有多少小于当前数字的数字
//https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number/
func main() {
	arr := []int{8, 1, 2, 2, 3}
	r := smallerNumbersThanCurrent(arr)
	fmt.Printf("%v", r)
}
func smallerNumbersThanCurrent(nums []int) []int {

	res := []int{}
	for _, v := range nums {
		temp := 0
		for _, v1 := range nums {
			if v1 < v {
				temp++
			}
		}
		res = append(res, temp)
	}
	return res

}
