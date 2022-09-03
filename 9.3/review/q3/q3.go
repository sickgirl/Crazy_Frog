package main

import "sort"

//查询最富有的用户
//https://leetcode.cn/problems/richest-customer-wealth/
func main() {
	account := [][]int{{1, 2, 3}, {1, 2, 3}}
	r := getRich(account)
	println(r)
}

func getRich(account [][]int) int {
	var rs []int
	for _, v1 := range account {
		temp := 0
		for _, v2 := range v1 {
			temp += v2
		}
		rs = append(rs, temp)
	}
	sort.Ints(rs)
	return rs[len(rs)-1]

}
